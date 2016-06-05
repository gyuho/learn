#!/bin/bash

go test -run=xxx -bench . -benchmem -cpu 1,2,4,8
