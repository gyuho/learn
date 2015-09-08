[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# Python: os, io

- [read file](#read-file)
- [write file](#write-file)
- [write csv](#write-csv)

[↑ top](#python-os-io)
<br><br><br><br>
<hr>








#### read file

```python
#!/usr/bin/python -u

if __name__ == "__main__":

    f = open('../README.md')
    print f.read()

```

[↑ top](#python-os-io)
<br><br><br><br>
<hr>









#### write file

```python
#!/usr/bin/python -u
import os

if __name__ == "__main__":

    fpath = "test.txt"

    f = open(fpath, 'w')
    f.write("Hello World!")
    f.close()

    f = open(fpath)
    print f.read()
    # Hello World!

    os.remove(fpath)

```

[↑ top](#python-os-io)
<br><br><br><br>
<hr>








#### write csv

```python
#!/usr/bin/python -u
import csv
import os

if __name__ == "__main__":

    fpath = "test.csv"
    columns = ["A", "B", "C"]

    data1 = {}
    data1["A"] = "Google"
    data1["B"] = "Facebook"
    data1["C"] = "Amazon"

    data2 = {}
    data2["A"] = "Google"
    data2["B"] = "Microsoft"
    data2["C"] = "Amazon"
    
    data3 = {}
    data3["A"] = "Amazon"
    data3["B"] = "Facebook"
    data3["C"] = "Google"
    
    data_list = [data1, data2, data3]

    with open(fpath, 'wb') as f:
        w = csv.DictWriter(f, fieldnames=columns, delimiter=',')
        w.writeheader()
        for elem in data_list:
            w.writerow(elem)

    os.remove(fpath)

```

[↑ top](#python-os-io)
<br><br><br><br>
<hr>
