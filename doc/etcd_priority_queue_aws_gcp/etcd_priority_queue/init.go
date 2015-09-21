package main

import (
	"database/sql"
	"flag"
	"fmt"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/coreos/etcd/client"
)

type (
	key int

	job struct {
		ETCDKey   string  `json:"etcd_key,omitempty"`
		ETCDIndex int     `json:"etcd_index,omitempty"`
		Priority  float64 `json:"priority,omitempty"`
		Action    string  `json:"action"`
		Status    string  `json:"status,omitempty"`
		Done      bool    `json:"done"`
	}
)

const (
	isInVPC = false

	odbKey key = 0
	xdbKey key = 1
	rdbKey key = 2
	cdbKey key = 3
)

var (
	panicCount   int
	recoverLimit = 100

	logPath = fmt.Sprintf("etcd_server_%s.log", strings.Replace(nowPacific().String()[:10], "-", "", -1))

	odb *sql.DB
	xdb *sql.DB
	rdb *sql.DB
	cdb *sql.DB

	dbTimeout    = 15 * time.Second
	queryTimeout = 5 * time.Second

	machines = []string{
		"http://1.2.3.0:2379", // etcd0
		"http://1.2.3.1:2379", // etcd1
		"http://1.2.3.2:2379", // etcd2
	}

	etcdClientConfig = client.Config{
		Endpoints: machines,
		Transport: client.DefaultTransport,
	}
)

func init() {
	log.SetFormatter(new(log.JSONFormatter))
	log.SetLevel(log.DebugLevel)
	logPathPtr := flag.String(
		"log",
		logPath,
		"Specify the log path.",
	)
	flag.Parse()
	lp := *logPathPtr
	if lp != logPath {
		logPath = lp
	}
}
