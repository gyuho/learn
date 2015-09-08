[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# Python: pointer, memory

- [reference](#reference)

[↑ top](#python-pointer-memory)
<br><br><br><br>
<hr>








#### reference

```python
#!/usr/bin/python -u

def change(tdict, tlist):
    tdict["A"] = True
    tlist.append(tdict)

if __name__ == "__main__":
    td = {}
    td["A"] = 100
    tl = [1,2,3]
    change(td, tl)
    print td # {'A': True}
    print tl # [1, 2, 3, {'A': True}]

```

For more, please read this
[article](https://www.jeffknupp.com/blog/2012/11/13/is-python-callbyvalue-or-callbyreference-neither/).

[↑ top](#python-pointer-memory)
<br><br><br><br>
<hr>
