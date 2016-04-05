#!/usr/bin/env bash
set -e

<<COMMENT
sudo su
nano /etc/sudoers
gyuho ALL=(ALL) NOPASSWD: ALL
COMMENT

##########################################################

##########################################################

sudo apt-get -y --force-yes update
sudo apt-get -y --force-yes upgrade
sudo apt-get -y --force-yes install build-essential
sudo apt-get -y --force-yes install git
sudo apt-get -y --force-yes install gcc
sudo apt-get -y --force-yes install bash curl git tar iptables iproute2 unzip
sudo apt-get -y --force-yes install bash-completion
sudo apt-get -y --force-yes install unzip gzip tar
sudo apt-get -y --force-yes install tree htop
sudo apt-get -y --force-yes install openssh
sudo apt-get -y --force-yes update
sudo apt-get -y --force-yes upgrade
sudo apt-get -y --force-yes autoremove
sudo apt-get -y --force-yes autoclean
echo "echo 1 > /proc/sys/vm/drop_caches" | sudo sh

sudo apt-get -y --force-yes install ntpdate
sudo service ntp stop
sudo ntpdate time.nist.gov
sudo service ntp start

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

go get -v -u -f github.com/tools/godep && \
go get -v -u -f golang.org/x/tools/cmd/... && \
go get -v -u -f github.com/golang/lint/golint && \
go get -v -u -f github.com/nsf/gocode && \
go get -v -u -f github.com/motain/gocheck && \
go get -v -u -f github.com/vaughan0/go-ini && \
go get -v -u -f github.com/rogpeppe/godef && \
go get -v -u -f github.com/kisielk/errcheck && \
go get -v -u -f github.com/jstemmer/gotags && \
go get -v -u -f github.com/alecthomas/gometalinter && \
go get -v -u -f golang.org/x/tools/cmd/benchcmp && \
go get -v -u -f golang.org/x/tools/cmd/goimports

cd $GOPATH/src/github.com/nsf/gocode/vim && sudo ./update.sh

##########################################################

##########################################################

sudo apt-get -y install vim && \
sudo apt-get -y install vim-nox && \
sudo apt-get -y install vim-gtk && \
sudo apt-get -y install vim-gnome && \
sudo apt-get -y install vim-athena

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

sudo apt-get -y install dh-autoreconf

cd $HOME/go/src/github.com && rm -rf google/protobuf && mkdir google
cd $HOME/go/src/github.com/google && git clone https://github.com/google/protobuf.git
cd $HOME/go/src/github.com/google/protobuf && ./autogen.sh
cd $HOME/go/src/github.com/google/protobuf && ./configure
cd $HOME/go/src/github.com/google/protobuf && make
cd $HOME/go/src/github.com/google/protobuf && make check;
cd $HOME/go/src/github.com/google/protobuf && make install

##########################################################
