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
sudo apt-get -y --force-yes install git
sudo apt-get -y --force-yes install gcc bash curl git tar iptables iproute2 unzip ntpdate bash-completion
sudo apt-get -y --force-yes install unzip gzip tar tree htop openssh
sudo apt-get -y --force-yes install dh-autoreconf
sudo apt-get -y --force-yes install vim vim-nox vim-gtk vim-gnome vim-athena

sudo apt-get -y --force-yes install tmux

echo "deb http://debian.sur5r.net/i3/ $(lsb_release -c -s) universe" >> /etc/apt/sources.list
sudo apt-get -y --force-yes update
sudo apt-get -y --force-yes --allow-unauthenticated install sur5r-keyring
sudo apt-get -y --force-yes update
sudo apt-get -y --force-yes install i3

sudo apt-get -y --force-yes remove unity unity-asset-pool unity-control-center unity-control-center-signon unity-gtk-module-common unity-lens* unity-services unity-settings-daemon unity-webapps* unity-voice-service
sudo apt-get -y --force-yes install lubuntu-desktop lxde

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

cd $HOME/go/src/github.com && rm -rf google/protobuf && mkdir google
cd $HOME/go/src/github.com/google && git clone https://github.com/google/protobuf.git
cd $HOME/go/src/github.com/google/protobuf && ./autogen.sh
cd $HOME/go/src/github.com/google/protobuf && ./configure
cd $HOME/go/src/github.com/google/protobuf && make
cd $HOME/go/src/github.com/google/protobuf && make check
cd $HOME/go/src/github.com/google/protobuf && make install
protoc --version

##########################################################
