[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# Python: sort

- [`sort`](#sort)
- [sort list of dictionaries](#sort-list-of-dictionaries)

[↑ top](#python-sort)
<br><br><br><br><hr>


#### `sort`

```python
#!/usr/bin/python -u

if __name__ == "__main__":
    li = [5, 4, 3, "A", "B"]
    li.sort()
    print li
    # [3, 4, 5, 'A', 'B']
```

[↑ top](#python-sort)
<br><br><br><br><hr>


#### sort list of dictionaries

```python
#!/usr/bin/python -u
import operator

if __name__ == "__main__":

    data1 = {}
    data1['date'] = "2015-08-01"
    data1['value'] = 100

    data2 = {}
    data2['date'] = "2019-08-01"
    data2['value'] = 500

    data3 = {}
    data3['date'] = "1900-08-01"
    data3['value'] = 900

    data_list = [data1, data2, data3]
    data_list.sort()

    new_data_list = sorted(data_list, key=operator.itemgetter('date'), reverse=True)
    print new_data_list
    # [{'date': '2019-08-01', 'value': 500}, {'date': '2015-08-01', 'value': 100}, {'date': '1900-08-01', 'value': 900}]

```

[↑ top](#python-sort)
<br><br><br><br><hr>
