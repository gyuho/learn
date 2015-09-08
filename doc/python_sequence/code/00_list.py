#!/usr/bin/python -u

if __name__ == "__main__":
    li = []
    li.append("A")
    li.append(1)
    li.append([1,2,3])
    li = li + ["X", "Y"]
    print li
    # ['A', 1, [1, 2, 3], 'X', 'Y']