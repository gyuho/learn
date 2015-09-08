package main

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func myPostgresDB() (*sql.DB, error) {
	return postgres("host=MY_HOST port=MY_PORT user=MY_USER password=MY_PASSWORD dbname=MY_DATABASE_NAME connect_timeout=2 sslmode=require")
}

func myMySQLDB() (*sql.DB, error) {
	return mysql("MY_USER:MY_PASSWORD@tcp(MY_HOST:MY_PORT)/MY_DATABASE_NAME?timeout=2s")
}

func main() {
	db, err := mustConn("myPostgresDB")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := mustQuery(db, "myPostgresDB", `select col1, col2 from schema.table`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var val1, val2 []byte
		if err := rows.Scan(
			&val1,
			&val2,
		); err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(val1))
		fmt.Println(string(val2))
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func postgres(conf string) (*sql.DB, error) {
	db, err := sql.Open("postgres", conf)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func mysql(conf string) (*sql.DB, error) {
	db, err := sql.Open("mysql", conf)
	if err != nil {
		return nil, err
	}
	return db, nil
}

var funcMap = map[string]func() (*sql.DB, error){
	"myMySQLDB":    myMySQLDB,
	"myPostgresDB": myPostgresDB,
}

const maxRetry = 3

func mustConn(opt string) (*sql.DB, error) {
	log.Printf("opening %s\n", opt)
	var funcToRun func() (*sql.DB, error)
	if v, ok := funcMap[opt]; ok {
		funcToRun = v
	} else {
		return nil, fmt.Errorf("mustConn Can't find the option for %s", opt)
	}
	var errMsg error
	for i := 0; i < maxRetry; i++ {
		if db, err := funcToRun(); err != nil {
			errMsg = err
			log.Printf("mustConn %s error:\n\n%s\n\n\n", opt, err)
			if i < 2 {
				log.Println("Waiting for 10 seconds... and retry...")
				time.Sleep(10 * time.Second)
			}
			db, err = funcToRun()
			if err != nil {
				errMsg = err
				log.Printf("mustConn %s error:\n\n%s\n\n\n", opt, err)
			}
			duration := time.Duration(float64ToInt64(math.Pow(2.0, float64(i))) * 100 * int64(time.Millisecond))
			log.Printf("Sleeping %+v and retrying...\n", duration)
			time.Sleep(duration)
			continue
		} else {
			errMsg = nil
			db.SetMaxOpenConns(100)
			return db, nil
		}
	}
	log.Printf("mustConn %s error:\n\n%s\n\n\n", opt, errMsg)
	return nil, fmt.Errorf("Failed after max-try: %d / %+v", maxRetry, errMsg)
}

func mustQuery(sdb *sql.DB, opt, query string) (*sql.Rows, error) {
	var sqlRows *sql.Rows
	var errMsg error
	for i := 0; i < maxRetry; i++ {
		if rows, err := sdb.Query(query); err != nil {
			errMsg = err
			log.Printf("mustQuery %s error:\n\n%s\n\n\n", opt, query)
			log.Println("Query error:", err)
			if i < 2 {
				log.Println("Waiting for 10 seconds... and retry...")
				time.Sleep(10 * time.Second)
			}
			sdb, err = mustConn(opt)
			if err != nil {
				errMsg = err
				log.Printf("mustQuery %s error:\n\n%s\n\n\n", opt, query)
				log.Println("Conn error:", err)
			}
			duration := time.Duration(float64ToInt64(math.Pow(2.0, float64(i))) * 100 * int64(time.Millisecond))
			log.Printf("Sleeping %+v and retrying...\n", duration)
			time.Sleep(duration)
			continue
		} else {
			sqlRows = rows
			errMsg = nil
			break
		}
	}
	if errMsg != nil {
		log.Printf("mustQuery %s error:\n\n%s\n\n\n", opt, query)
		return nil, fmt.Errorf("Failed after max-try: %d", maxRetry)
	}
	return sqlRows, nil
}

func mustExec(sdb *sql.DB, opt, query string) error {
	var errMsg error
	for i := 0; i < maxRetry; i++ {
		if _, err := sdb.Exec(query); err != nil {
			errMsg = err
			log.Printf("mustExec %s error:\n\n%s\n\n\n", opt, query)
			log.Println("Query error:", err)
			if i < 2 {
				log.Println("Waiting for 10 seconds... and retry...")
				time.Sleep(10 * time.Second)
			}
			sdb, err = mustConn(opt)
			if err != nil {
				errMsg = err
				log.Printf("mustExec %s error:\n\n%s\n\n\n", opt, query)
				log.Println("Conn error:", err)
			}
			duration := time.Duration(float64ToInt64(math.Pow(2.0, float64(i))) * 100 * int64(time.Millisecond))
			log.Printf("Sleeping %+v and retrying...\n", duration)
			time.Sleep(duration)
			continue
		} else {
			errMsg = nil
			break
		}
	}
	if errMsg != nil {
		log.Printf("mustExec %s error:\n\n%s\n\n\n", opt, query)
		return fmt.Errorf("Failed after max-try: %d", maxRetry)
	}
	return nil
}

func mustExecTransaction(sdb *sql.DB, queries ...string) error {
	tx, err := sdb.Begin()
	if err != nil {
		return err
	}

	for _, query := range queries {
		log.Println("mustExecTransaction Executing:", query)
		var errMsg error
		for i := 0; i < maxRetry; i++ {
			if _, err := tx.Exec(query); err != nil {
				errMsg = err
				log.Printf("mustExecTransaction error:\n\n%s\n\n\n", query)
				log.Println("Query error:", err)
				if i < 2 {
					log.Println("Waiting for 10 seconds... and retry...")
					time.Sleep(10 * time.Second)
				}

				duration := time.Duration(float64ToInt64(math.Pow(2.0, float64(i))) * 100 * int64(time.Millisecond))
				log.Printf("Sleeping %+v and retrying...\n", duration)
				time.Sleep(duration)
				continue
			} else {
				errMsg = nil
				break
			}
		}
		if errMsg != nil {
			log.Printf("mustExecTransaction error:\n\n%s\n\n\n", query)
			return fmt.Errorf("Failed after max-try: %d", maxRetry)
		}
	}

	var errMsg error
	for i := 0; i < maxRetry; i++ {
		if err := tx.Commit(); err != nil {
			errMsg = err
			log.Printf("Commit error:\n\n%s\n\n\n", errMsg)
			log.Println("Commit error:", err)
			if i < 2 {
				log.Println("Waiting for 10 seconds... and retry...")
				time.Sleep(10 * time.Second)
			}

			duration := time.Duration(float64ToInt64(math.Pow(2.0, float64(i))) * 100 * int64(time.Millisecond))
			log.Printf("Sleeping %+v and retrying...\n", duration)
			time.Sleep(duration)
			continue
		} else {
			errMsg = nil
			break
		}
	}
	if errMsg != nil {
		log.Printf("mustExecTransaction error:\n\n%s\n\n\n", errMsg)
		return fmt.Errorf("Failed after max-try: %d", maxRetry)
	}

	return nil
}

func float64ToInt64(num float64) int64 {
	return int64(num)
}
