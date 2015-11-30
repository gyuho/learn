package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	ct1 := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	ct2 := parseTS("2009-11-10 23:00:00")
	if !ct1.Equal(ct2) {
		panic(fmt.Errorf("expected \n%v\n%v", ct1, ct2))
	}
	fmt.Println(ct2)
	fmt.Println(parseDate0("2016-07-07"))
	fmt.Println(parseDate1("November 7, 2017"))
	fmt.Println(parseDate2("06/07/2019"))
	fmt.Println(substractDate("2016-07-07", "2015-07-07"))
	/*
	   2009-11-10 23:00:00 +0000 UTC
	   2016-07-07 00:00:00 +0000 UTC
	   2017-11-07 00:00:00 -0800 PST
	   2019-06-07 00:00:00 -0700 PDT
	   366
	*/
}

// parseDate0 parses the string-format date time.
func parseDate0(stamp string) time.Time {
	rt, err := time.Parse("2006-01-02", stamp)
	if err != nil {
		panic(err)
	}
	return rt
}

func parseDate1(date string) time.Time {
	zoneSC, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}
	pt, err := time.Parse("January 2, 2006", date)
	if err != nil {
		panic(err)
	}
	rd := time.Date(
		pt.Year(),
		pt.Month(),
		pt.Day(),
		0,
		0,
		0,
		0,
		zoneSC)
	return rd
}

func parseDate2(date string) time.Time {
	zoneSC, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}
	pt, err := time.Parse("01/02/2006", date)
	if err != nil {
		panic(err)
	}
	rd := time.Date(
		pt.Year(),
		pt.Month(),
		pt.Day(),
		0,
		0,
		0,
		0,
		zoneSC)
	return rd
}

// substractDate returns the difference between two dates.
// (dt1 - dt2)
func substractDate(dt1, dt2 string) int {
	td1 := parseDate0(dt1)
	td2 := parseDate0(dt2)
	diff := td1.Sub(td2).Hours()
	return int(diff / 24)
}

// parseTS parses the string-format time stamp.
func parseTS(stamp string) time.Time {
	stamp = strings.TrimSpace(stamp)
	if len(strings.Split(stamp, ".")) > 1 {
		mc1 := strings.Split(stamp, ".")[0]
		mc2 := strings.Split(stamp, ".")[1]
		if len(mc2) == 1 {
			stamp = mc1 + "." + mc2 + "00000"
		}
		if len(mc2) == 2 {
			stamp = mc1 + "." + mc2 + "0000"
		}
		if len(mc2) == 3 {
			stamp = mc1 + "." + mc2 + "000"
		}
		if len(mc2) == 4 {
			stamp = mc1 + "." + mc2 + "00"
		}
		if len(mc2) == 5 {
			stamp = mc1 + "." + mc2 + "0"
		}
	} else if len(strings.Split(stamp, ".")) == 1 {
		stamp = stamp + ".000000"
	}

	rt, err := time.Parse("2006-01-02 15:04:05.000000", stamp)
	if err != nil {
		panic(err)
	}
	return rt
}
