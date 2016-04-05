#!/usr/bin/env bash
set -e

<<COMMENT
sudo su
nano /etc/sudoers
gyuho ALL=(ALL) NOPASSWD: ALL
COMMENT

##########################################################

##########################################################
# Ubuntu

sudo apt-get -y --force-yes update
sudo apt-get -y --force-yes upgrade

sudo apt-get -y --force-yes install build-essential 
sudo apt-get -y --force-yes install git mercurial meld
sudo apt-get -y --force-yes install gcc bash curl git tar iptables iproute2 unzip ntpdate bash-completion unzip gzip tar tree htop
sudo apt-get -y --force-yes install dh-autoreconf xclip autoconf automake libtool
sudo apt-get -y --force-yes install vim vim-nox vim-gtk vim-gnome vim-athena terminator

sudo apt-get -y --force-yes update
sudo apt-get -y --force-yes upgrade
sudo apt-get -y --force-yes autoremove
sudo apt-get -y --force-yes autoclean

sudo service ntp stop
sudo ntpdate time.nist.gov
sudo service ntp start

echo "echo 1 > /proc/sys/vm/drop_caches" | sudo sh

##########################################################

##########################################################

echo "[user]
  email = gyuhox@gmail.com
  name = Gyu-Ho Lee

[color]
  diff = auto
  status = auto
  branch = auto
  ui = auto" > $HOME/.gitconfig

git config --global user.name "Gyu-Ho Lee"
git config --global user.email "gyuhox@gmail.com"
git config --global core.editor "vim"

ssh-keygen -t rsa -b 4096 -C "gyuhox@gmail.com"
eval "$(ssh-agent -s)"
ssh-add $HOME/.ssh/id_rsa
xclip -sel clip < $HOME/.ssh/id_rsa.pub

##########################################################

##########################################################

cd $HOME && sudo rm -rf go1.4		
cd $HOME && sudo rm -rf go1.4_temp && mkdir -p $HOME/go1.4_temp		
sudo curl -s https://storage.googleapis.com/golang/go1.4.linux-amd64.tar.gz | sudo tar -v -C $HOME/go1.4_temp -xz		
cd $HOME/go1.4_temp && sudo mv ./go ./go1.4 && sudo mv ./go1.4 ..
cd $HOME && sudo rm -rf $HOME/go1.4_temp

GO_VERSION="1.6" && cd /usr/local && sudo rm -rf ./go && sudo curl -s https://storage.googleapis.com/golang/go$GO_VERSION.linux-amd64.tar.gz | sudo tar -v -C /usr/local/ -xz && cd $HOME;
PATH_VAR=$PATH":/usr/local/go/bin:$(echo $HOME)/go/bin"
export GOPATH=$(echo $HOME)/go
PATH_VAR=$PATH":/usr/local/go/bin:$(echo $HOME)/go/bin"
export PATH=$(echo $PATH_VAR)
go version

if grep -q GOPATH "$(echo $HOME)/.bashrc"; then 
	echo "bashrc already has GOPATH...";
else
	echo "adding GOPATH to bashrc...";
	echo "export GOPATH=$(echo $HOME)/go" >> $HOME/.bashrc && \
	PATH_VAR=$PATH":/usr/local/go/bin:$(echo $HOME)/go/bin" && \
	echo "export PATH=$(echo $PATH_VAR)" >> $HOME/.bashrc && \
	source $HOME/.bashrc;
fi
# echo "export PATH=$(echo $PATH_VAR)" >> $HOME/.bashrc
# echo "export GOPATH=$(echo $HOME)/go" >> $HOME/.bashrc

cd $HOME && rm -rf ./go-master && mkdir -p $HOME/go-master && \
cd $HOME/go-master && git clone https://go.googlesource.com/go && cd $HOME/go-master/go/src && ./all.bash

##########################################################

##########################################################

mkdir -p $HOME/go/src/github.com/gyuho
cd $HOME/go/src/github.com/gyuho
git clone git@github.com:gyuho/learn.git

mkdir -p $HOME/go/src/github.com/coreos
cd $HOME/go/src/github.com/coreos
git clone git@github.com:coreos/etcd.git

go get -v -u -f github.com/tools/godep && \
go get -v -u -f github.com/golang/lint/golint && \
go get -v -u -f github.com/nsf/gocode && \
go get -v -u -f github.com/motain/gocheck && \
go get -v -u -f github.com/vaughan0/go-ini && \
go get -v -u -f github.com/rogpeppe/godef && \
go get -v -u -f github.com/kisielk/errcheck && \
go get -v -u -f github.com/jstemmer/gotags && \
go get -v -u -f github.com/alecthomas/gometalinter && \
go get -v -u -f golang.org/x/tools/cmd/benchcmp && \
go get -v -u -f golang.org/x/tools/cmd/goimports && \
go get -v -u -f golang.org/x/tools/cmd/vet && \
go get -v -u -f golang.org/x/tools/cmd/oracle && \
go get -v -u -f honnef.co/go/simple/cmd/gosimple && \
go get -v -u -f honnef.co/go/unused/cmd/unused

cd $GOPATH/src/github.com/nsf/gocode/vim && sudo ./update.sh

go get -v -u -f github.com/gyuho/psn && \
go get -v -u -f github.com/gyuho/gomp && \
go get -v -u -f github.com/coreos/dbtester && \
go get -v -u -f github.com/coreos/etcd-play && \
go get -v -u -f github.com/coreos/etcd

##########################################################

##########################################################

sudo chown -R gyuho:gyuho $HOME/.vim
sudo mkdir -p $HOME/.vim/ftdetect
sudo mkdir -p $HOME/.vim/syntax
sudo chmod -R +x $HOME/.vim

curl -fLo $HOME/.vim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

<<COMMENT
:PlugInstall
:PlugClean
:PlugUpdate
:PlugUpgrade
:GoInstallBinaries
COMMENT

##########################################################

##########################################################

PROTOC_VERSION=3.0.0-beta-2
curl -sf -o /tmp/protoc.zip https://github.com/google/protobuf/releases/download/v$PROTOC_VERSION/protoc-$PROTOC_VERSION-linux-x86_64.zip
unzip /tmp/protoc.zip -d /usr/bin/
rm -f /tmp/protoc.zip
protoc --version

##########################################################
