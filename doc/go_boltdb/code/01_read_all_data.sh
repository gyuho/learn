#!/bin/bash

echo "echo 1 > /proc/sys/vm/drop_caches" | sudo sh;
go run 01_read_all_data.go -populate=true;

printf "\n"

echo "echo 1 > /proc/sys/vm/drop_caches" | sudo sh;
go run 01_read_all_data.go -populate=false;
