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