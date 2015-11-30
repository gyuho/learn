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
