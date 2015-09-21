[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# Go: introduction

- [Reference](#reference)
- [Install](#install)
- [Hello World!](#hello-world)

[↑ top](#go-introduction)
<br><br><br><br>
<hr>







#### Reference

- [golang.org](http://golang.org/)
- [Go Tour](http://tour.golang.org/welcome/1)

[↑ top](#go-introduction)
<br><br><br><br>
<hr>









#### Install

Please visit [here](http://golang.org/doc/install).
In `Ubuntu 14.04.3 LTS` (Linux Debian distribution), you can run:

```bash
#!/bin/bash

cd $HOME;
mkdir -p $HOME/go/src/github.com;
mkdir -p $HOME/go/src/golang.org;
sudo curl -s https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz | sudo tar -v -C /usr/local/ -xz;

echo "export GOPATH=$(echo $HOME)/go" >> $HOME/.bashrc;
TEMP_PATH=$PATH':/usr/local/go/bin:/home/ubuntu/go/bin'
echo "export PATH=$(echo $TEMP_PATH)" >> $HOME/.bashrc;
source $HOME/.bashrc;

cd $HOME;
printf "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Successfully installed Go\")\n}" > $HOME/temp.go; 
go run temp.go; 
rm -rf temp.go;
go version;

go get -v -u github.com/tools/godep;
go get -v -u golang.org/x/tools/cmd/goimports;
go get -v -u github.com/golang/lint/golint;
go get -v -u github.com/nsf/gocode;
go get -v -u github.com/motain/gocheck;
go get -v -u github.com/vaughan0/go-ini;
cd $GOPATH/src/github.com/nsf/gocode/vim; sudo ./update.sh;

cd $HOME;
go get -v -u github.com/bradrydzewski/go.auth;
go get -v -u github.com/coreos/etcd/...;
go get -v -u github.com/coreos/etcd/...;
go get -v -u github.com/go-sql-driver/mysql;
go get -v -u github.com/gyuho/awsapi/redshift;
go get -v -u github.com/gyuho/cloudflare;
go get -v -u github.com/gyuho/htmlx;
go get -v -u github.com/jordan-wright/email;
go get -v -u github.com/lib/pq;
go get -v -u github.com/satori/go.uuid;
go get -v -u github.com/Sirupsen/logrus;
go get -v -u github.com/Sirupsen/logrus;
go get -v -u github.com/tylerb/graceful;
go get -v -u golang.org/x/net/context;
go get -v -u golang.org/x/net/context;
go get -v -u gopkg.in/yaml.v2;

cd $HOME;

```

[↑ top](#go-introduction)
<br><br><br><br>
<hr>









#### Hello World!

Try this [code](http://play.golang.org/p/OccSs5jC9Y):

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

You can either:

- `go run hello/main.go`
- `cd hello/` and `go build` and `./hello`
- `cd hello/` and `go install` and `hello`

[↑ top](#go-introduction)
<br><br><br><br>
<hr>
