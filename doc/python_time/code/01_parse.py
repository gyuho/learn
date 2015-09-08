#!/usr/bin/python -u

import datetime

if __name__ == "__main__":
    dt_str = '9/24/2015 7:00:00 PM'
    dt_obj = datetime.datetime.strptime(dt_str, '%m/%d/%Y %I:%M:%S %p')
    print dt_obj, type(dt_obj)
    # 2015-09-24 19:00:00 <type 'datetime.datetime'>
