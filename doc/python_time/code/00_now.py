#!/usr/bin/python -u
import datetime
import pytz

PST_TZ = pytz.timezone("US/Pacific")
DAY = datetime.timedelta(hours=24)

def pacific_now():
    return datetime.datetime.now(PST_TZ).replace(tzinfo=None)

def pacific_yesterday():
    return pacific_now() - DAY

def utc_now():
    return datetime.datetime.utcnow()

def datetime_to_date_str(dt):
    return '%04d-%02d-%02d' % (dt.year, dt.month, dt.day)

def datetime_to_timestamp_str(dt):
    return dt.strftime("%Y-%m-%d %H:%M:%S")

if __name__ == "__main__":
    print pacific_now()
    print pacific_yesterday()
    print utc_now()
    print datetime_to_date_str(pacific_now())      # 2015-08-13
    print datetime_to_timestamp_str(pacific_now()) # 2015-08-13 20:53:07

    print datetime.datetime.utcnow() - datetime.timedelta(days=2)
    # 2015-08-12 03:55:09.357782
