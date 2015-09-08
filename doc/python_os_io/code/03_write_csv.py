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
