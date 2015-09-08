[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# Python: time

- [Now with timezone](#now-with-timezone)
- [Parse](#parse)
- [Calendar](#calendar)

[↑ top](#python-time)
<br><br><br><br>
<hr>








#### Now with timezone

```python
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

```

[↑ top](#python-time)
<br><br><br><br>
<hr>









#### Parse

```python
#!/usr/bin/python -u

import datetime

if __name__ == "__main__":
    dt_str = '9/24/2015 7:00:00 PM'
    dt_obj = datetime.datetime.strptime(dt_str, '%m/%d/%Y %I:%M:%S %p')
    print dt_obj, type(dt_obj)
    # 2015-09-24 19:00:00 <type 'datetime.datetime'>

```

[↑ top](#python-time)
<br><br><br><br>
<hr>









#### Calendar

```python
#!/usr/bin/python -u
import datetime
import calendar

if __name__ == "__main__":
    print datetime.datetime.today().weekday() # 3
    print calendar.day_name[datetime.datetime.today().weekday()]
    # Thursday

```

[↑ top](#python-time)
<br><br><br><br>
<hr>
