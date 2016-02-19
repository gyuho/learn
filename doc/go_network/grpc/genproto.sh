#!/usr/bin/env bash
#
# Generate all etcd protobuf bindings.
# Run from repository root.
#
set -e

# for now, be conservative about what version of protoc we expect
if ! [[ $(protoc --version) =~ "3.0.0" ]]; then
	echo "could not find protoc 3.0.0, is it installed + in PATH?"
	exit 255
fi

#!/usr/bin/env bash
go get -v -u -f github.com/gogo/protobuf/proto;
go get -v -u -f github.com/gogo/protobuf/protoc-gen-gogo;
go get -v -u -f github.com/gogo/protobuf/gogoproto;
go get -v -u -f github.com/gogo/protobuf/protoc-gen-gofast;

protoc \
	--gofast_out=plugins=grpc:. \
	--proto_path=$GOPATH/src:$GOPATH/src/github.com/gogo/protobuf/protobuf:. \
	*.proto;
