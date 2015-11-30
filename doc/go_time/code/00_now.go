package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println(nowPST())
	fmt.Println(nowUTC())
	fmt.Println(todayPST())
	fmt.Println(todayUTC())
	fmt.Println(dateTag(nowPST()))
	fmt.Println(timeTag(nowPST()))
	/*
	   2015-08-09 21:17:50.508866464 -0700 PDT
	   2015-08-10 04:17:50.508917948 +0000 UTC
	   2015-08-09 00:00:00 -0700 PDT
	   2015-08-10 00:00:00 +0000 UTC
	   20150809
	   2015080921175050897
	*/
}

// Manipulate Time with:
// func (t Time) Add(d Duration) Time
// func (t Time) AddDate(years int, months int, days int) Time

func nowPST() time.Time {
	tzone, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Returning time.Now()")
		return time.Now()
	}
	return time.Now().In(tzone)
}

func nowUTC() time.Time {
	tzone, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Returning time.Now()")
		return time.Now()
	}
	return time.Now().In(tzone).UTC()
}

func todayPST() time.Time {
	tzone, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Returning time.Now()")
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

func todayUTC() time.Time {
	tzone, err := time.LoadLocation("")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Returning time.Now()")
		return time.Date(
			time.Now().Year(),
			time.Now().Month(),
			time.Now().Day(),
			0,
			0,
			0,
			0,
			time.UTC,
		)
	}
	utc := time.Now().In(tzone).UTC()
	return time.Date(
		utc.Year(),
		utc.Month(),
		utc.Day(),
		0,
		0,
		0,
		0,
		tzone)
}

func stringByScale(t time.Time, scale string) string {
	switch scale {
	case "micro":
		// 2014-06-23 15:37:21.12311
		return t.String()[:25]
	case "second":
		return t.String()[:19]
	case "date":
		return t.String()[:10]
	default:
		return t.String()
	}
}

func dateTag(t time.Time) string {
	ts := stringByScale(t, "date")
	return strings.Replace(ts, "-", "", -1)
}

func timeTag(t time.Time) string {
	ts := stringByScale(t, "micro")
	ts = strings.Replace(ts, "-", "", -1)
	ts = strings.Replace(ts, ":", "", -1)
	ts = strings.Replace(ts, ".", "", -1)
	ts = strings.Replace(ts, " ", "", -1)
	return ts
}
