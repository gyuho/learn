#!/usr/bin/python -u
print type(1), type("A"), type([1,2,3]), type({"A":True}), type(True)
# <type 'int'> <type 'str'> <type 'list'> <type 'dict'> <type 'bool'>

print isinstance(1, int)
# True

print isinstance("A", str)
# True

