#!/usr/bin/python -u
import datetime
import calendar

if __name__ == "__main__":
    print datetime.datetime.today().weekday() # 3
    print calendar.day_name[datetime.datetime.today().weekday()]
    # Thursday
