#!/bin/bash
go get -v github.com/dvyukov/go-fuzz/go-fuzz;
go get -v github.com/dvyukov/go-fuzz/go-fuzz-build;

go-fuzz-build github.com/coreos/etcd/raft;
go-fuzz -bin=./raft-fuzz.zip -workdir=raft;

go-fuzz-build github.com/coreos/etcd/rafthttp;
go-fuzz -bin=./rafthttp-fuzz.zip -workdir=rafthttp;

go-fuzz-build github.com/coreos/etcd/storage;
go-fuzz -bin=./storage-fuzz.zip -workdir=storage;
