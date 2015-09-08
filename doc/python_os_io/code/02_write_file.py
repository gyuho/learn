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
