[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# Python: sequence

- [list](#list)
- [immutable tuple](#immutable-tuple)

[↑ top](#python-sequence)
<br><br><br><br>
<hr>





#### list

```python
#!/usr/bin/python -u

if __name__ == "__main__":
    li = []
    li.append("A")
    li.append(1)
    li.append([1,2,3])
    li = li + ["X", "Y"]
    print li
    # ['A', 1, [1, 2, 3], 'X', 'Y']
```

[↑ top](#python-sequence)
<br><br><br><br>
<hr>









#### immutable tuple

```python
#!/usr/bin/python -u

if __name__ == "__main__":
    ti = (1,2,3)
    print ti
    # (1, 2, 3)

```

[↑ top](#python-sequence)
<br><br><br><br>
<hr>
