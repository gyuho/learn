#!/bin/bash

printf "\n"
echo "TEST #1 $ go test -v ./..."
go test -v ./...;

printf "\n"
echo "TEST #2 $ go test -v -race ./.."
go test -v -race ./...;

printf "\n"
echo "TEST #3 $ go test -bench . -benchmem -cpu 1,2,4,8,16"
go test -opt="Slice" -bench . -benchmem -cpu 1,2,4,8,16 > slice.txt;
go test -opt="Map"   -bench . -benchmem -cpu 1,2,4,8,16 > map.txt;

go get -v -u golang.org/x/tools/cmd/benchcmp;
benchcmp slice.txt map.txt > benchmark_results.txt;

printf "\n"
echo "Done"
