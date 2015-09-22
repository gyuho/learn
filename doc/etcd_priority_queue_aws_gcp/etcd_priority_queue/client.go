package main

import (
	"bytes"
	"container/heap"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/net/context"

	log "github.com/Sirupsen/logrus"
	"github.com/coreos/etcd/client"
)

/*
curl http://$INFRA_PUBLIC_IP_0:2379/v2/keys/queue
curl http://$INFRA_PUBLIC_IP_0:2379/v2/keys/queue/{EtcdIndex}

curl -L -XPUT http://$INFRA_PUBLIC_IP_0:2379/v2/keys/queue/{EtcdIndex} -d value="Hello"
curl -L -XPOST http://$INFRA_PUBLIC_IP_0:2379/v2/keys/queue/{EtcdIndex} -d value="Hello"
curl -L -XDELETE http://$INFRA_PUBLIC_IP_0/v2/keys/queue/{EtcdIndex}
*/

const dirName = "queue"

func getJobs(ctx context.Context, kapi client.KeysAPI) (map[string]*job, error) {
	rmapChan := make(chan map[string]*job)
	errChan := make(chan error)
	go func() {
		opts := &client.GetOptions{}
		opts.Recursive = false
		opts.Sort = false
		resp, err := kapi.Get(ctx, dirName, opts)
		if err != nil {
			errChan <- err
			return
		}
		if resp == nil {
			errChan <- fmt.Errorf("Empty Response: %+v", resp)
			return
		}
		if resp.Node == nil {
			fmt.Printf("Empty Queue: %+v\n", resp)
			rmapChan <- nil
			return
		}
		if resp.Node.Nodes.Len() == 0 {
			fmt.Printf("Empty Queue: %+v\n", resp)
			rmapChan <- nil
			return
		}
		queueMap := make(map[string]*job)
		for _, elem := range resp.Node.Nodes {
			if _, ok := queueMap[elem.Key]; !ok {
				j := job{}
				if err := json.NewDecoder(strings.NewReader(elem.Value)).Decode(&j); err != nil {

					log.WithFields(log.Fields{
						"event_type": "error",
						"value":      elem.Value,
						"error":      err,
					}).Errorln("getJobs json Decode error")

					j.Done = true
					j.Status = err.Error()
					if err := setJob(ctx, kapi, &j); err != nil {
						errChan <- err
						return
					}
					continue
				}
				j.ETCDKey = elem.Key
				id := strings.Replace(elem.Key, "/"+dirName+"/", "", -1)
				iv, err := strconv.Atoi(id)
				if err != nil {
					log.WithFields(log.Fields{
						"event_type": "error",
						"error":      err,
					}).Errorln("getJobs strconv.Atoi error")
					j.Done = true
					j.Status = err.Error()
					if err := setJob(ctx, kapi, &j); err != nil {
						errChan <- err
						return
					}
					continue
				}
				j.ETCDIndex = iv
				queueMap[elem.Key] = &j
				rmapChan <- queueMap
				return
			}
		}
	}()
	select {
	case v := <-rmapChan:
		return v, nil

	case v := <-errChan:
		return nil, v

	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func setJob(ctx context.Context, kapi client.KeysAPI, j *job) error {
	errChan := make(chan error)
	done := make(chan struct{})
	go func() {
		buf := new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(j); err != nil {
			errChan <- err
			return
		}
		value := buf.String()
		fmt.Println("setJob:", value)
		opts := &client.SetOptions{}
		if _, err := kapi.Set(ctx, j.ETCDKey, value, opts); err != nil {
			errChan <- err
			return
		}
		done <- struct{}{}
	}()
	select {
	case <-done:
		return nil

	case v := <-errChan:
		return v

	case <-ctx.Done():
		return ctx.Err()
	}
}

func createJobInOrder(ctx context.Context, kapi client.KeysAPI, j *job) error {
	errChan := make(chan error)
	done := make(chan struct{})
	go func() {
		buf := new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(j); err != nil {
			errChan <- err
			return
		}
		value := buf.String()
		fmt.Println("createJobInOrder:", value)
		opts := &client.CreateInOrderOptions{}
		if _, err := kapi.CreateInOrder(ctx, dirName, value, opts); err != nil {
			errChan <- err
			return
		}
		done <- struct{}{}
	}()
	select {
	case <-done:
		return nil

	case v := <-errChan:
		return v

	case <-ctx.Done():
		return ctx.Err()
	}
}

func runJobs(ctx context.Context, kapi client.KeysAPI) error {
	errChan := make(chan error)
	done := make(chan struct{})
	go func() {
		jmap, err := getJobs(ctx, kapi)
		if err != nil {
			errChan <- err
			return
		}
		var pq priorityQueue
		for _, v := range jmap {
			if v == nil {
				continue
			}
			if v.Done {
				continue
			}
			heap.Push(&pq, v)
		}
		heap.Init(&pq)
		if pq.Len() == 0 {
			fmt.Println("priorityQueue is empty")
			done <- struct{}{}
			return
		}

		println()
		fmt.Println("# of jobs:", pq.Len())
		println()
		xd := ctx.Value(xdbKey).(*sql.DB)

		for pq.Len() > 0 {
			j := heap.Pop(&pq).(*job)
			log.WithFields(log.Fields{
				"event_type": "runJobs",
				"etcd_key":   j.ETCDKey,
				"etcd_index": j.ETCDIndex,
				"action":     j.Action,
			}).Debugln("runJobs")

			rt, cancel := context.WithTimeout(context.Background(), queryTimeout)
			defer cancel()
			rows, err := runQuery(
				rt,
				xd,
				fmt.Sprintf(`SELECT value FROM table`),
			)
			if err != nil {
				errChan <- err
				return
			}
			defer rows.Close()
			for rows.Next() {
				var v []byte
				if err := rows.Scan(&v); err != nil {
					log.WithFields(log.Fields{
						"event_type": "rows.Scan",
						"error":      err,
					}).Errorln("runJobs")
					break
				}
				_ = string(v)
				break
			}

			switch j.Action {
			case "doSomething":
				done <- struct{}{}
				return
			default:
				errChan <- fmt.Errorf("unknown action: %+v", *j)
				return
			}
		}
	}()
	select {
	case <-done:
		fmt.Println("runJobs done.")
		return nil

	case v := <-errChan:
		return v

	case <-ctx.Done():
		return ctx.Err()
	}
}
