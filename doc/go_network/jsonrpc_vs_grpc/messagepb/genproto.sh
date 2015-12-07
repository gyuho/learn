#!/usr/bin/env bash
go get -v -u github.com/gogo/protobuf/proto;
go get -v -u github.com/gogo/protobuf/protoc-gen-gogo;
go get -v -u github.com/gogo/protobuf/gogoproto;
go get -v -u github.com/gogo/protobuf/protoc-gen-gofast;

protoc \
	--gofast_out=plugins=grpc:. \
	--proto_path=$GOPATH/src:$GOPATH/src/github.com/gogo/protobuf/protobuf:. \
	*.proto;
