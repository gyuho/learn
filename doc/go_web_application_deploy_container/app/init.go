package main

import (
	"database/sql"
	"flag"
	"fmt"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
)

type (
	key int

	storage struct {
		sync.Mutex
		userIDToData map[string]*data
	}

	data struct {
		Sentence string
	}

	Output struct {
		Name  string
		Value float64
		TS    string
	}
)

const (
	port = ":8080"

	isInVPC = false

	// this is a test key in my test account.
	googleClientID     = "883943096730-780g2kk9pinusp94qgm6orrad5qt316v.apps.googleusercontent.com"
	googleClientSecret = "I4rws53mkbPp0288EpdBGzyT"

	userKey key = 0
	odbKey  key = 1
	xdbKey  key = 2
	rdbKey  key = 3
	cdbKey  key = 4
)

var (
	panicCount   int
	recoverLimit = 10

	logPath = fmt.Sprintf("%s.log", nowPacific().String()[:10])

	odb *sql.DB
	xdb *sql.DB
	rdb *sql.DB
	cdb *sql.DB

	dbTimeout    = 15 * time.Second
	queryTimeout = 5 * time.Second

	globalStorage storage

	accessibleEmail = map[string]bool{
		"gyuhox@gmail.com": true,
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
