#!/bin/bash
go get -v github.com/bradrydzewski/go.auth;
go get -v github.com/coreos/etcd/...;
go get -v github.com/go-sql-driver/mysql;
go get -v github.com/go-yaml/yaml;
go get -v github.com/gyuho/htmlx;
go get -v github.com/jordan-wright/email;
go get -v github.com/lib/pq;
go get -v github.com/satori/go.uuid;
go get -v github.com/Sirupsen/logrus;
go get -v github.com/tylerb/graceful;
go get -v golang.org/x/net/context;

sudo kill $(sudo netstat -tlpn | perl -ne 'my @a = split /[ \/]+/; print "$a[6]\n" if m/:8080/gio');
./app 1>>stdout.log 2>>stderr.log;
