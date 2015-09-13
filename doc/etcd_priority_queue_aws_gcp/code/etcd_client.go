package main

import (
	"bytes"
	"container/heap"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

var machines = []string{
	"http://1.2.3.4:2379", // etcd0
	"http://1.2.3.5:2379", // etcd1
	"http://1.2.3.6:2379", // etcd2
}

var cfg = client.Config{
	Endpoints: machines,
	Transport: client.DefaultTransport,
	// set timeout per request to fail fast when the target endpoint is unavailable
	HeaderTimeoutPerRequest: time.Second,
}

func main() {
	recursiveRun()
}

var count int

// panic and recover
func recursiveRun() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered but here's the error:", err)
			count++
			if count == 2 {
				panic("panic for this time. Debug this!")
			}
			run()
		}
	}()
	run()
}

func run() {
	totalCount := 0
	now := time.Now()

	for {
		totalCount++

		c, err := client.New(cfg)
		if err != nil {
			panic(err)
		}
		kapi := client.NewKeysAPI(c)

		if err := execute(kapi, "queue"); err != nil {
			fmt.Println("error:", err)
			fmt.Println("running duration:", time.Since(now))
			fmt.Println("total count:", totalCount)
			panic(err)
		}

		time.Sleep(5 * time.Second)
	}
}

// Job wraps generate API command request and response.
// This has to be mapped from etcd index.
type Job struct {
	ETCDKey string `json:"etcd_key,omitempty"`

	// max-heap returns the element with the highest priority
	// where we compare the Priority first and then ETCDIndex.
	ETCDIndex int     `json:"etcd_index,omitempty"`
	Priority  float64 `json:"priority,omitempty"`

	Action string `json:"action"`
	Status string `json:"status,omitempty"`
	Done   bool   `json:"done"`
}

// get retrieves all the jobs in an etcd server.
// It only gets the Jobs that have not been done.
//
//   curl http://$INFRA_PUBLIC_IP_0:2379/v2/keys/queue
//
func get(kapi client.KeysAPI, queueName string) (map[string]*Job, error) {
	resp, err := kapi.Get(context.Background(), ctx, queueName, nil)
	if err != nil {
		if err == context.Canceled {
			return nil, fmt.Errorf("ctx is canceled by another routine")
		} else if err == context.DeadlineExceeded {
			return nil, fmt.Errorf("ctx is attached with a deadline and it exceeded")
		} else if cerr, ok := err.(*client.ClusterError); ok {
			return nil, fmt.Errorf("*client.ClusterError %v", cerr.Errors())
		} else {
			return nil, fmt.Errorf("bad cluster endpoints, which are not etcd servers: %+v", err)
		}
	}
	if resp == nil {
		log.Printf("Empty resp: %+v\n", resp)
		return nil, nil
	}
	if resp.Node == nil {
		log.Printf("Empty Queue: %+v\n", resp)
		return nil, nil
	}
	if resp.Node.Nodes.Len() == 0 {
		log.Printf("Empty Queue: %+v\n", resp)
		return nil, nil
	}
	queueMap := make(map[string]*Job)
	for _, elem := range resp.Node.Nodes {
		if _, ok := queueMap[elem.Key]; !ok {
			job := Job{}
			if err := json.NewDecoder(strings.NewReader(elem.Value)).Decode(&job); err != nil {
				log.Println("json.NewDecoder error:", elem.Value, err)
				continue
			}
			job.ETCDKey = elem.Key
			ids := strings.Replace(elem.Key, "/"+queueName+"/", "", -1)
			idv, err := strconv.Atoi(ids)
			if err != nil {
				log.Println("strconv.Atoi error:", ids, err)
				continue
			}
			job.ETCDIndex = idv
			queueMap[elem.Key] = &job
		}
	}
	return queueMap, nil
}

// put updates/creates the job.
func put(kapi client.KeysAPI, job *Job) error {
	job.FinishedTimestamp = time.Now().UTC().String()[:19]
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(job); err != nil {
		return err
	}
	value := buf.String()
	if _, err := kapi.Set(context.Background(), ctx, job.ETCDKey, value, nil); err != nil {
		if err == context.Canceled {
			return fmt.Errorf("ctx is canceled by another routine")
		} else if err == context.DeadlineExceeded {
			return fmt.Errorf("ctx is attached with a deadline and it exceeded")
		} else if cerr, ok := err.(*client.ClusterError); ok {
			return fmt.Errorf("*client.ClusterError %v", cerr.Errors())
		} else {
			return fmt.Errorf("bad cluster endpoints, which are not etcd servers: %+v", err)
		}
	}
	return nil
}

// priorityQueue is a min-heap of Jobs.
type priorityQueue []*Job

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	p1 := pq[i].Priority
	idx1 := pq[i].ETCDIndex

	p2 := pq[j].Priority
	idx2 := pq[j].ETCDIndex

	if p1 == p2 {
		// min-heap returns the lowest priority first
		// when the Priority's were same, we want to return the one with lower index.
		return idx1 < idx2
	}

	// max-heap
	return p1 > p2
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Job))
}

func (pq *priorityQueue) Pop() interface{} {
	heapSize := len(*pq)
	lastNode := (*pq)[heapSize-1]
	*pq = (*pq)[:heapSize-1]
	return lastNode
}

func execute(kapi *client.KeysAPI, queueName string) error {
	log.Println("#1. get")
	jobMap, err := get(kapi, queueName)
	if err != nil {
		return err
	}

	log.Println("#2. Max-Heapify all jobs")
	var pq priorityQueue
	for _, job := range jobMap {
		heap.Push(&pq, job)
	}

	// heapify
	heap.Init(&pq)

	log.Println("We have", pq.Len(), "elements in the queue")

	log.Println("#3. Pop, Execute, put one by one, until the heap is empty")
	for pq.Len() > 0 {
		oneJob := heap.Pop(&pq).(*Job)
		if oneJob == nil {
			continue
		}

		if oneJob.Done {
			log.Printf("Done, so this is skipped: %+v", oneJob)
			continue
		}

		if pq.Len() == 0 {
			log.Println("Executing the last element")
		}

		// executing the job
		fmt.Println(oneJob.Action)
	}

	log.Println("Done with queue...")
}
