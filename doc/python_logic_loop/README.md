[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# Python: logic, loop

- [logic](#logic)
- [if](#if)
- [loop](#loop)

[↑ top](#python-logic-loop)
<br><br><br><br><hr>


#### logic

```python
#!/usr/bin/python -u

if __name__ == "__main__":
    print True or False  # True
    print True and False # False
    print False or False # False

    print (1 == 1) or ([] == "A")
    # True

    print ([] == "A") or (1 == 2)
    # False 
```

[↑ top](#python-logic-loop)
<br><br><br><br><hr>


#### if

```python
#!/usr/bin/python -u

if __name__ == "__main__":

    if 1 == 2:
        print "1 == 2"
    elif 2 == 3:
        print "2 == 3"
    else:
        print "None"

    # None
```

[↑ top](#python-logic-loop)
<br><br><br><br><hr>


#### loop

```python
#!/usr/bin/python -u

if __name__ == "__main__":
    
    array = [0,1,2,3,4,5]
    for k in array:
        print k

    """
    0
    1
    2
    3
    4
    5
    """

    print

    td = {"A": True, "B": False, "C": True}
    for k in td:
        print k, td[k]
    for k, v in td.iteritems():
        print k, v

    """
    A True
    C True
    B False
    A True
    C True
    B False
    """

    cnt = 0
    while True:

        print "Hello World!"

        cnt += 1
        if cnt == 5:
            break
        else:
            continue

        print "THIS IS NOT PRINTED"

    """
    Hello World!
    Hello World!
    Hello World!
    Hello World!
    Hello World!
    """
```

[↑ top](#python-logic-loop)
<br><br><br><br><hr>
