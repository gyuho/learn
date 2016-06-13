#!/bin/bash

go test -v -benchmem -run=xxx -bench=BenchmarkStringSwitch
go test -v -benchmem -run=xxx -bench=BenchmarkStringMap
