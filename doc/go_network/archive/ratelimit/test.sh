#!/usr/bin/env bash
# Reference: 
# https://github.com/coreos/etcd/blob/master/test

TEST=./...;
FMT="*.go"

echo "Running tests...";
go test -v $TEST;
go test -v -race $TEST;

echo "Checking gofmt..."
fmtRes=$(gofmt -l -s $FMT)
if [ -n "${fmtRes}" ]; then
	echo -e "gofmt checking failed:\n${fmtRes}"
	exit 255
fi

echo "Checking govet..."
vetRes=$(go vet $TEST)
if [ -n "${vetRes}" ]; then
	echo -e "govet checking failed:\n${vetRes}"
	exit 255
fi

echo "Success";

