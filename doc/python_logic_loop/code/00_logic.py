#!/usr/bin/python -u

if __name__ == "__main__":
    print True or False  # True
    print True and False # False
    print False or False # False

    print (1 == 1) or ([] == "A")
    # True

    print ([] == "A") or (1 == 2)
    # False 