[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# Go: time

- [Now with timezone](#now-with-timezone)
- [Parse](#parse)
- [Calendar](#calendar)

[↑ top](#go-time)
<br><br><br><br><hr>


#### Now with timezone

```go
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
```

[↑ top](#go-time)
<br><br><br><br><hr>


#### Parse

```go
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
```

[↑ top](#go-time)
<br><br><br><br><hr>


#### Calendar

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ct := time.Date(2014, time.July, 10, 07, 0, 0, 0, time.UTC)
	if weekDay(ct) != "Thu" {
		panic(fmt.Errorf("expected Thu but %v", weekDay(ct)))
	}

	ct1 := time.Date(2014, time.July, 11, 07, 0, 0, 0, time.UTC)
	str1, bl1 := isWeekend(ct1)
	if str1 != "Fri" || bl1 != false {
		panic(fmt.Errorf("expected %v\n%v", str1, bl1))
	}

	ct2 := time.Date(2014, time.July, 12, 03, 0, 0, 0, time.UTC)
	str2, bl2 := isWeekend(ct2)
	if str2 != "Sat" || bl2 != true {
		panic(fmt.Errorf("expected %v\n%v", str2, bl2))
	}

	ct3 := time.Date(2014, time.July, 13, 17, 0, 0, 0, time.UTC)
	str3, bl3 := isWeekend(ct3)
	if str3 != "Sun" || bl3 != true {
		panic(fmt.Errorf("expected %v\n%v", str3, bl3))
	}

	newyearsday := time.Date(2014, time.January, 01, 07, 0, 0, 0, time.UTC)
	if str, bl := isUSHoliday(newyearsday); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	martinluther := time.Date(2014, time.January, 20, 07, 0, 0, 0, time.UTC)
	if str, bl := isUSHoliday(martinluther); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	valentine := time.Date(2014, time.February, 14, 07, 0, 0, 0, time.UTC)
	if str, bl := isUSHoliday(valentine); bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	if str, bl := isHalloweenValentine(valentine); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	washington := time.Date(2014, time.February, 17, 07, 0, 0, 0, time.UTC)
	if str, bl := isUSHoliday(washington); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	mother := time.Date(2014, time.May, 11, 07, 0, 0, 0, time.UTC)
	if str, bl := isMotherFather(mother); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	memorial := time.Date(2014, time.May, 26, 07, 0, 0, 0, time.UTC)
	if str, bl := isUSHoliday(memorial); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	father := time.Date(2014, time.June, 15, 07, 0, 0, 0, time.UTC)
	if str, bl := isMotherFather(father); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	independence := time.Date(2014, time.July, 4, 07, 0, 0, 0, time.UTC)
	if str, bl := isUSHoliday(independence); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	labor := time.Date(2014, time.September, 1, 07, 0, 0, 0, time.UTC)
	if str, bl := isUSHoliday(labor); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	columbus := time.Date(2014, time.October, 13, 07, 0, 0, 0, time.UTC)
	if str, bl := isUSHoliday(columbus); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	halloween := time.Date(2014, time.October, 31, 07, 0, 0, 0, time.UTC)
	if str, bl := isHalloweenValentine(halloween); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	veteran := time.Date(2014, time.November, 11, 07, 0, 0, 0, time.UTC)
	if str, bl := isUSHoliday(veteran); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	thanksgiving := time.Date(2014, time.November, 27, 07, 0, 0, 0, time.UTC)
	if str, bl := isUSHoliday(thanksgiving); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	christmas := time.Date(2014, time.December, 25, 07, 0, 0, 0, time.UTC)
	if str, bl := isUSHoliday(christmas); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	christmaseve := time.Date(2014, time.December, 24, 07, 0, 0, 0, time.UTC)
	if str, bl := isSemiHoliday(christmaseve); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}

	newyeareve := time.Date(2014, time.December, 31, 07, 0, 0, 0, time.UTC)
	if str, bl := isSemiHoliday(newyeareve); !bl {
		panic(fmt.Errorf("expected %v, %v", str, bl))
	}
}

func weekDay(t time.Time) string {
	const layout = "Mon"
	return t.Format(layout)
}

func isWeekend(t time.Time) (string, bool) {
	day := weekDay(t)
	if day == "Sat" || day == "Sun" {
		return day, true
	}
	return day, false
}

func isUSHoliday(t time.Time) (string, bool) {
	switch t.Month() {
	case time.January:
		if t.Day() == 1 {
			return "New Year's Day", true
		}
		// third Monday
		if weekDay(t) == "Mon" &&
			t.AddDate(0, 0, -7*2).Month() == time.January &&
			t.AddDate(0, 0, -7*3).Month() == time.December {
			return "Martin Luther King, Jr. Day", true
		}
	case time.February:
		// third Monday
		if weekDay(t) == "Mon" &&
			t.AddDate(0, 0, -7*2).Month() == time.February &&
			t.AddDate(0, 0, -7*3).Month() == time.January {
			return "Washington's Birthday", true
		}
	case time.May:
		// last Monday
		if weekDay(t) == "Mon" &&
			t.AddDate(0, 0, 7*1).Month() == time.June {
			return "Memorial Day", true
		}
	case time.July:
		if t.Day() == 4 {
			return "Independence Day", true
		}
	case time.September:
		// first Monday
		if weekDay(t) == "Mon" &&
			t.AddDate(0, 0, -7*1).Month() == time.August {
			return "Labor Day", true
		}
	case time.October:
		// second Monday
		if weekDay(t) == "Mon" &&
			t.AddDate(0, 0, -7*1).Month() == time.October &&
			t.AddDate(0, 0, -7*2).Month() == time.September {
			return "Columbus Day", true
		}
	case time.November:
		if t.Day() == 11 {
			return "Veterans Day", true
		}
		// fourth Thursday
		if weekDay(t) == "Thu" &&
			t.AddDate(0, 0, -7*3).Month() == time.November &&
			t.AddDate(0, 0, -7*4).Month() == time.October {
			return "Thanksgiving Day", true
		}
	case time.December:
		if t.Day() == 25 {
			return "Christmas Day", true
		}
	}
	return "None", false
}

// isSemiHoliday returns true if it is American Holiday Eve,
// or the day after Thanksgiving day.
func isSemiHoliday(t time.Time) (string, bool) {
	// the day after Thanksgiving Day
	st, _ := isUSHoliday(t.AddDate(0, 0, -1))
	if st == "Thanksgiving Day" {
		return "Day After Thanksgiving Day", true
	}
	tm := t.AddDate(0, 0, 1)
	sh, bl := isUSHoliday(tm)
	return "Before " + sh, bl
}

// isHalloweenValentine returns true if it is Halloween or Valentine's Day.
func isHalloweenValentine(t time.Time) (string, bool) {
	if t.Month() == time.October {
		if t.Day() == 31 {
			return "Halloween", true
		}
	}
	if t.Month() == time.February {
		if t.Day() == 14 {
			return "Valentine's Day", true
		}
	}
	return "None", false
}

// isMotherFather returns true if it is Mother's day or Father's day.
func isMotherFather(t time.Time) (string, bool) {
	if t.Month() == time.May {
		// second Sunday
		if weekDay(t) == "Sun" &&
			t.AddDate(0, 0, -7*1).Month() == time.May &&
			t.AddDate(0, 0, -7*2).Month() == time.April {
			return "Mother's Day", true
		}
	}
	if t.Month() == time.June {
		// third Sunday
		if weekDay(t) == "Sun" &&
			t.AddDate(0, 0, -7*2).Month() == time.June &&
			t.AddDate(0, 0, -7*3).Month() == time.May {
			return "Father's Day", true
		}
	}
	return "None", false
}
```

[↑ top](#go-time)
<br><br><br><br><hr>

