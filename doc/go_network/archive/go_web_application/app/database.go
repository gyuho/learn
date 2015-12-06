package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"golang.org/x/net/context"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var funcMap = map[string]func(bool) (*sql.DB, error){
	"xx": xx,
	"oo": oo,
	"rr": rr,
	"cc": cc,
}

func postgres(cfg string) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func mysql(cfg string) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func xx(isInVPC bool) (*sql.DB, error) {
	if isInVPC {
		return postgres("host=private_endpoint port=5439 user=ID password=PASSWORD dbname=xx connect_timeout=2 sslmode=require")
	}
	return postgres("host=public_endpoint port=5439 user=ID password=PASSWORD dbname=omdw connect_timeout=2 sslmode=require")
}

func oo(isInVPC bool) (*sql.DB, error) {
	if isInVPC {
		return mysql("ID:PASSWORD@tcp(pricate_endpoint:3306)/cc?timeout=2s")
	}
	return mysql("ID:PASSWORD@tcp(public_endpoint:3306)/cc?timeout=2s")
}

func rr(isInVPC bool) (*sql.DB, error) {
	if isInVPC {
		return postgres("host=private_endpoint port=5432 user=ID password=PASSWORD dbname=rr connect_timeout=2 sslmode=require")
	}
	return postgres("host=public_endpoint port=5432 user=ID password=PASSWORD dbname=report connect_timeout=2 sslmode=require")
}

func cc(isInVPC bool) (*sql.DB, error) {
	if isInVPC {
		return mysql("ID:PASSWORD@tcp(pricate_endpoint:3306)/cc?timeout=2s")
	}
	return mysql("ID:PASSWORD@tcp(public_endpoint:3306)/cc?timeout=2s")
}

func getConn(ctx context.Context, isInVPC bool, dbname string) (*sql.DB, error) {
	log.Println("opening", dbname)
	var funcToRun func(bool) (*sql.DB, error)
	if v, ok := funcMap[dbname]; ok {
		funcToRun = v
	} else {
		return nil, fmt.Errorf("conn can't find the dbname %s", dbname)
	}
	done := make(chan struct{})
	var (
		db     *sql.DB
		errMsg error
	)
	go func() {
		for {
			d, err := funcToRun(isInVPC)
			if err != nil {
				errMsg = err
				time.Sleep(time.Second)
				continue
			} else {
				db = d
				errMsg = nil
				done <- struct{}{}
				break
			}
		}
	}()
	select {
	case <-done:
		log.Println("opened", dbname)
		return db, errMsg
	case <-ctx.Done():
		return nil, fmt.Errorf("getConn %s timed out with %v / %v", dbname, ctx.Err(), errMsg)
	}
}

func runQuery(ctx context.Context, db *sql.DB, query string) (*sql.Rows, error) {
	done := make(chan struct{})
	var (
		rows   *sql.Rows
		errMsg error
	)
	go func() {
		for {
			rs, err := db.Query(query)
			if err != nil {
				errMsg = err
				time.Sleep(time.Second)
				continue
			} else {
				rows = rs
				errMsg = nil
				done <- struct{}{}
				break
			}
		}
	}()
	select {
	case <-done:
		return rows, errMsg
	case <-ctx.Done():
		return nil, fmt.Errorf("runQuery %s timed out with %v / %v", query, ctx.Err(), errMsg)
	}
}

func runExec(ctx context.Context, db *sql.DB, query string) error {
	done := make(chan struct{})
	var (
		errMsg error
	)
	go func() {
		for {
			if _, err := db.Exec(query); err != nil {
				errMsg = err
				time.Sleep(time.Second)
				continue
			} else {
				errMsg = nil
				done <- struct{}{}
				break
			}
		}
	}()
	select {
	case <-done:
		return errMsg
	case <-ctx.Done():
		return fmt.Errorf("runExec %s timed out with %v / %v", query, ctx.Err(), errMsg)
	}
}
