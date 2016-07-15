#!/usr/bin/env bash
set -e

# for now, be conservative about what version of protoc we expect
if ! [[ $(protoc --version) =~ "3.0.0" ]]; then
	echo "could not find protoc 3.0.0, is it installed + in PATH?"
	exit 255
fi

echo "Installing golang/protobuf..."
GOLANG_PROTO_ROOT="$GOPATH/src/github.com/golang/protobuf"
rm -rf $GOLANG_PROTO_ROOT
go get -v -u github.com/golang/protobuf/{proto,protoc-gen-go} 
go get -v -u golang.org/x/tools/cmd/goimports
pushd "${GOLANG_PROTO_ROOT}"
	git reset --hard HEAD
	make install
popd

printf "Generating proto\n"
protoc --go_out=plugins=grpc:. \
	*.proto;

go get -v -u github.com/golang/protobuf/jsonpb

