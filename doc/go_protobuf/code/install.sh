#!/bin/bash
sudo apt-get install protobuf-compiler;
go get -v github.com/golang/protobuf/{proto,protoc-gen-go};

cd /home/ubuntu/go/src/github.com/gyuho/learn/doc/go_protobuf/example;
protoc --go_out=. *.proto;
# output: example.pb.go

cd /home/ubuntu/go/src/github.com/gyuho/learn/doc/go_protobuf/code;
go run main.go;
