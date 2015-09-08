#!/usr/bin/python -u
import random

if __name__ == "__main__":
	random.seed(random.randint(0, 999))
	rv = random.uniform(-0.5, 0.5)
	print rv
	# 0.456034271889
