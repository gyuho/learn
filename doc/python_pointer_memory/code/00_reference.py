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
