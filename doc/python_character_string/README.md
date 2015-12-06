[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# Python: character, string

- [immutable string](#immutable-string)
- [escape special characters](#escape-special-characters)
- [`string` literals in Python](#string-literals-in-python)

[↑ top](#python-character-string)
<br><br><br><br><hr>


#### immutable string

```python
#!/usr/bin/python -u

if __name__ == "__main__":
    txt = "Hello"
    for c in txt:
        print c
    """
    H
    e
    l
    l
    o
    """

    # txt[1] = "X"
    # print txt
    # TypeError: 'str' object does not support item assignment

```

[↑ top](#python-character-string)
<br><br><br><br><hr>


#### escape special characters

```python
#!/usr/bin/python -u

if __name__ == "__main__":
    print "\\"  # \
    print "%%"  # %
    print "\""  # "
```

[↑ top](#python-character-string)
<br><br><br><br><hr>


#### `string` literals in Python

As you see the code below, Python has different string literals for `ASCII` and
`Unicode` characters.

```python
val1 = "aaé"
print val1        # aaé
print type(val1)  # <type 'str'>
 
print val1.encode('utf-8')
"""
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
UnicodeDecodeError: 'ascii' codec can't decode byte 0xc3 in position 0: ordinal not in range(128)
"""
 
print val1.encode('ascii')
"""
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
UnicodeDecodeError: 'ascii' codec can't decode byte 0xc3 in position 0: ordinal not in range(128)
"""
 
val2 = u"aaé"
print val2                  # aaé
print type(val2)            # <type 'unicode'>
print val2.encode('utf-8')  # aaé
 
print val2.encode('ascii')
"""
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
UnicodeEncodeError: 'ascii' codec can't encode character u'\xe9' in position 0: ordinal not in range(128)
"""
 
 
print val2.encode('ascii', 'ignore') # aa
# é is missing
 
 
import unicodedata
unicodedata.normalize('NFKD', val2).encode('ascii','ignore')
# aae
# é got converted to e
```

This can be tricky when an external service returns **different types** of
string to your Python program, as described
[here](https://pythonhosted.org/kitchen/unicode-frustrations.html):

> Frustration #1: Inconsistent Errors
>
> Although converting when possible seems like the right thing to do, it’s
> actually the first source of frustration. A programmer can test out their
> program with a string like: *The quick brown fox jumped over the lazy dog* and
> not encounter any issues. But when they release their software into the wild,
> someone enters the string: *I sat down for coffee at the café* and suddenly an
> exception is thrown. The reason? The mechanism that converts between the two
> types is only able to deal with
> [ASCII](https://pythonhosted.org/kitchen/glossary.html#term-ascii) characters. 
> Once you throw non-ASCII characters into your strings, you have to start
> dealing with the conversion manually.
>
> So, if I manually convert everything to either byte str or unicode strings,
> will I be okay? The answer is…. sometimes.
>
> [Overcoming frustration: Correctly using unicode in
> python2](https://pythonhosted.org/kitchen/unicode-frustrations.html)


Here’s a quick solution:

```python
def convert_to_str(st):
    """ use this function to convert all strings to str"""
    if isinstance(st, unicode):
        return st.encode('utf-8')
    return str(st)
 
 
val1 = "ébc"
val2 = u"ébc"
 
print val1, type(val1), convert_to_str(val1), type(convert_to_str(val1))
# ébc <type 'str'> ébc <type 'str'>
 
print val2, type(val2), convert_to_str(val2), type(convert_to_str(val2))
# ébc <type 'unicode'> ébc <type 'str'>
```

But there are many other corner cases you need to consider, as explained
[here](https://pythonhosted.org/kitchen/unicode-frustrations.html).

[↑ top](#python-character-string)
<br><br><br><br><hr>
