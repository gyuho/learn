package main

import (
	"time"

	"golang.org/x/net/context"

	stdlog "log"

	log "github.com/Sirupsen/logrus"
	"github.com/coreos/etcd/client"
)

func main() {
	keepRunning()
}

func keepRunning() {
	f, err := openToAppend(logPath)
	if err != nil {
		stdlog.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)

	cx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	xx, err := getConn(cx, isInVPC, "xx")
	if err != nil {
		log.Panic(err)
	}
	xdb = xx

	co, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	oo, err := getConn(co, isInVPC, "oo")
	if err != nil {
		log.Panic(err)
	}
	odb = oo

	cr, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	rr, err := getConn(cr, isInVPC, "rr")
	if err != nil {
		log.Panic(err)
	}
	rdb = rr

	cc, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	cd, err := getConn(cc, isInVPC, "cc")
	if err != nil {
		log.Panic(err)
	}
	cdb = cd

	defer odb.Close()
	defer xdb.Close()
	defer rdb.Close()
	defer cdb.Close()

	defer func() {
		if err := recover(); err != nil {

			log.WithFields(log.Fields{
				"event_type": "panic_recover",
				"error":      err,
			}).Errorln("keepRunning error")

			panicCount++
			if panicCount == recoverLimit {
				log.WithFields(log.Fields{
					"event_type": "panic",
					"error":      err,
				}).Panicln("Too much panic:", panicCount)
			}

			keepRunning()
		}
	}()

	mainRun()
}

func mainRun() {
	rootContext := context.Background()
	rootContext = context.WithValue(rootContext, odbKey, odb)
	rootContext = context.WithValue(rootContext, xdbKey, xdb)
	rootContext = context.WithValue(rootContext, rdbKey, rdb)
	rootContext = context.WithValue(rootContext, cdbKey, cdb)

	for {
		c, err := client.New(etcdClientConfig)
		if err != nil {
			panic(err)
		}
		kapi := client.NewKeysAPI(c)

		log.WithFields(log.Fields{
			"event_type": "run",
			"start_time": nowPacific(),
		}).Debugln("client.RunJobs running")

		if err := runJobs(rootContext, kapi); err != nil {
			log.WithFields(log.Fields{
				"event_type": "error",
				"error":      err,
			}).Errorln("client.RunJobs error")
		}

		time.Sleep(5 * time.Second)
	}
}
