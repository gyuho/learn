#!/bin/bash

x=$(lsb_release -a | grep "Distributor ID:")
if [ ${x:16:6} = "Ubuntu" ] || [ ${x:16:6} = "Debian" ]; then
    echo "Ubuntu ou debian"
    sudo apt-get -y install git
elif [ ${x:16:4} = "arch" ]; then
    echo "Arch linux"
    sudo pacman --noconfirm -S git
else
    echo "Distro unknown!"
fi

cd $HOME && \
mkdir -p $HOME/go/src && \
mkdir -p $HOME/go/src/github.com && \
mkdir -p $HOME/go/src/github.com/coreos && \
mkdir -p $HOME/go/src/github.com/gyuho && \
mkdir -p $HOME/go/src/golang.org;

cd /usr/local && sudo rm -rf ./go;

sudo curl \
-s https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz \
| sudo tar -v -C /usr/local/ -xz;

echo "export GOPATH=$(echo $HOME)/go" >> $HOME/.bashrc && \
PATH_VAR=$PATH":/usr/local/go/bin:$(echo $HOME)/go/bin" && \
echo "export PATH=$(echo $PATH_VAR)" >> $HOME/.bashrc && \
source $HOME/.bashrc;

cd $HOME && \
printf "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Successfully installed Go.\")\n}" > $HOME/temp.go;
cd $HOME && \
go run temp.go && \
rm -f temp.go && \
go version;

cd $HOME && \
go get -v -u github.com/tools/godep && \
go get -v -u golang.org/x/tools/cmd/... && \
go get -v -u github.com/golang/lint/golint && \
go get -v -u github.com/nsf/gocode && \
go get -v -u github.com/motain/gocheck && \
go get -v -u github.com/vaughan0/go-ini && \
go get -v -u github.com/rogpeppe/godef && \
go get -v -u github.com/kisielk/errcheck && \
go get -v -u github.com/jstemmer/gotags && \
cd $GOPATH/src/github.com/nsf/gocode/vim && sudo ./update.sh && \
cd $HOME;
