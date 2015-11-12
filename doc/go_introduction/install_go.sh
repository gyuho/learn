#!/bin/bash

cd $HOME;
mkdir -p $HOME/go/src/github.com;
mkdir -p $HOME/go/src/golang.org;

printf "\n\n\nDeleting GOROOT.\n" && \
cd /usr/local && sudo rm -rf ./go;

printf "\n\n\nInstalling Go.\n" && \
sudo curl -s https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz \
| sudo tar -v -C /usr/local/ -xz;

printf "\n\n\nConfiguring GOPATH.\n" && \
echo "export GOPATH=$(echo $HOME)/go" >> $HOME/.bashrc;
TEMP_PATH=$PATH':/usr/local/go/bin:/home/ubuntu/go/bin'
echo "export PATH=$(echo $TEMP_PATH)" >> $HOME/.bashrc;
source $HOME/.bashrc;

printf "\n\n\nTesting Go installation with a script.\n" && \
cd $HOME;
printf "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Successfully installed Go.\")\n}" > $HOME/temp.go; 
go run temp.go && rm -rf temp.go;

printf "\n\n\nPrinting Go version.\n" && go version;

printf "\n\n\nInstalling Go packages.\n"
go get -v -u github.com/tools/godep;
go get -v -u golang.org/x/tools/cmd/goimports;
go get -v -u github.com/golang/lint/golint;
go get -v -u github.com/nsf/gocode;
go get -v -u github.com/motain/gocheck;
go get -v -u github.com/vaughan0/go-ini;
go get -v -u github.com/rogpeppe/godef;
go get -v -u golang.org/x/tools/cmd/oracle;
go get -v -u golang.org/x/tools/cmd/gorename;
go get -v -u golang.org/x/tools/cmd/benchcmp;
go get -v -u golang.org/x/tools/cmd/...;
go get -v -u github.com/kisielk/errcheck;
go get -v -u github.com/jstemmer/gotags;

cd $GOPATH/src/github.com/nsf/gocode/vim && sudo ./update.sh;

cd $HOME;
go get -v -u github.com/bradrydzewski/go.auth;
go get -v -u github.com/go-sql-driver/mysql;
go get -v -u github.com/gyuho/awsapi/redshift;
go get -v -u github.com/gyuho/cloudflare;
go get -v -u github.com/gyuho/htmlx;
go get -v -u github.com/jordan-wright/email;
go get -v -u github.com/lib/pq;
go get -v -u github.com/satori/go.uuid;
go get -v -u github.com/Sirupsen/logrus;
go get -v -u github.com/tylerb/graceful;
go get -v -u golang.org/x/net/context;
go get -v -u gopkg.in/yaml.v2;
cd $HOME;

