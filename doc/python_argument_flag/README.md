[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# Python: argument, flag

- [flag](#flag)

[↑ top](#python-argument-flag)
<br><br><br><br><hr>


#### flag

```python
#!/usr/bin/python -u
import sys
import getopt

usage = """
Usage:
    ./00_flag.py -i 8 -o hello -d 7 -f

"""

try:
    opts, args = getopt.gnu_getopt(sys.argv[1:], "i:o:d:f")

except:
    print usage
    sys.exit(0)


if __name__ == "__main__":
    o = {}
    for opt, arg in opts:
        o[opt] = arg

    print '-f' in o

    from pprint import pprint
    pprint(o)

"""
./00_flag.py -i 8 -o hello -d 7 -f
True
{'-d': '7', '-f': '', '-i': '8', '-o': 'hello'}
"""

```

[↑ top](#python-argument-flag)
<br><br><br><br><hr>
