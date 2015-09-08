[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# Python: dictionary

- [dictionary](#dictionary)

[↑ top](#python-dictionary)
<br><br><br><br>
<hr>








#### dictionary

```python
#!/usr/bin/python -u

if __name__ == "__main__":
    td1 = {"A": True, "B": False, "G": True}
    for k in td1:
        print k, td1[k]

    td2 = {"A": True, "B": False, "G": True}
    for k in td2.keys():
        print k, td2[k]

    td3 = {"A": True, "B": False, "G": True}
    for k, v in td3.iteritems():
        print k, v

    print "----------------------"

    for k in td1:
        td1[k] = 100
    print td1
    # {'A': 100, 'B': 100, 'G': 100}

    # for k in td1:
    #     del td1[k]
    #     # RuntimeError: dictionary changed size during iteration
    # print td1

    print "----------------------"

    for k in td2:
        td2[k] = 100
    print td2
    # {'A': 100, 'B': 100, 'G': 100}

    for k in td2.keys():
        del td2[k]
    print td2
    # {}

    print "----------------------"

    for k, v in td3.iteritems():
        td3[k] = 100
    print td3
    # {'A': 100, 'B': 100, 'G': 100}

    # for k, v in td3.iteritems():
    #     del td3[k]
    #     # RuntimeError: dictionary changed size during iteration
    # print td3
```

[↑ top](#python-dictionary)
<br><br><br><br>
<hr>
