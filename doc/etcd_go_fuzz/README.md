[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# etcd, go-fuzz

This is to try break [`coreos/etcd`](https://github.com/coreos/etcd)
with [`dvyukov/go-fuzz`](https://github.com/dvyukov/go-fuzz):

- [Reference](#reference)
- [install](#install)
- [go-fuzz `etcdserver/server_test.go`](#go-fuzz-etcdserverserver_testgo)
- [go-fuzz `etcdserver/config_test.go`](#go-fuzz-etcdserverconfig_testgo)
- [go-fuzz `etcdserver/member_test.go`](#go-fuzz-etcdservermember_testgo)
- [go-fuzz `etcdserver/etcdhttp/client_test.go`](#go-fuzz-etcdserveretcdhttpclient_testgo)
- [go-fuzz `etcdserver/etcdhttp/client_auth_test.go`](#go-fuzz-etcdserveretcdhttpclient_auth_testgo)
- [go-fuzz `etcdserver/etcdhttp/peer_test.go`](#go-fuzz-etcdserveretcdhttppeer_testgo)
- [go-fuzz `storage/index_test.go`](#go-fuzz-storageindex_testgo)
- [go-fuzz `storage/kv_test.go`](#go-fuzz-storagekv_testgo)
- [go-fuzz `storage/key_index_test.go`](#go-fuzz-storagekey_index_testgo)

[↑ top](#etcd-go-fuzz)
<br><br><br><br>
<hr>








#### Reference

- [`dvyukov/go-fuzz`](https://github.com/dvyukov/go-fuzz)

[↑ top](#etcd-go-fuzz)
<br><br><br><br>
<hr>









#### install

```bash
#!/bin/bash
go get -v github.com/dvyukov/go-fuzz/go-fuzz;
go get -v github.com/dvyukov/go-fuzz/go-fuzz-build;
go get -v github.com/coreos/etcd/...;

```

[↑ top](#etcd-go-fuzz)
<br><br><br><br>
<hr>








#### test `etcdserver/server_test.go`



go-fuzz-build github.com/coreos/etcd/raft;
go-fuzz -bin=./raft-fuzz.zip -workdir=raft;

cp ~ to dir 

go-fuzz-build github.com/coreos/etcd/rafthttp;
go-fuzz -bin=./rafthttp-fuzz.zip -workdir=rafthttp;

go-fuzz-build github.com/coreos/etcd/storage;
go-fuzz -bin=./storage-fuzz.zip -workdir=storage;


[↑ top](#etcd-go-fuzz)
<br><br><br><br>
<hr>
