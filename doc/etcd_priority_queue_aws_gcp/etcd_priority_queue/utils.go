package main

import (
	"log"
	"os"
	"strings"
	"time"
)

func openToAppend(fpath string) (*os.File, error) {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		f, err = os.Create(fpath)
		if err != nil {
			return f, err
		}
	}
	return f, nil
}

func cleanUp(str string) string {
	s := strings.Fields(strings.TrimSpace(str))
	return strings.Join(s, " ")
}

func nowPacific() time.Time {
	tzone, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		return time.Now()
	}
	return time.Now().In(tzone)
}

func todayPacific() time.Time {
	tzone, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		log.Println(err)
		log.Println("Returning time.Now()")
		return time.Date(
			time.Now().Year(),
			time.Now().Month(),
			time.Now().Day(),
			0,
			0,
			0,
			0,
			nil,
		)
	}
	pst := time.Now().In(tzone)
	return time.Date(
		pst.Year(),
		pst.Month(),
		pst.Day(),
		0,
		0,
		0,
		0,
		tzone)
}

func weekday(t time.Time) string {
	const layout = "Mon"
	return t.Format(layout)
}
