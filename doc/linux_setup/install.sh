#!/usr/bin/env bash
set -e

<<COMMENT
sudo su
nano /etc/sudoers
gyuho ALL=(ALL) NOPASSWD: ALL
COMMENT

##########################################################

##########################################################

sudo apt-get -y --allow-unauthenticated update
sudo apt-get -y --allow-unauthenticated upgrade

sudo apt-get -y --allow-unauthenticated install build-essential
sudo apt-get -y --allow-unauthenticated install git mercurial meld
sudo apt-get -y --allow-unauthenticated install gcc bash curl git tar iptables iproute2 unzip ntpdate bash-completion unzip gzip tar tree htop
sudo apt-get -y --allow-unauthenticated install dh-autoreconf xclip autoconf automake libtool
sudo apt-get -y --allow-unauthenticated install vim vim-nox vim-gtk vim-gnome vim-athena ncurses-dev
sudo apt-get -y --allow-unauthenticated install terminator pcmanfm xclip
sudo apt-get -y --allow-unauthenticated install libpcap-dev libaspell-dev libhunspell-dev

sudo apt-get -y --allow-unauthenticated update
sudo apt-get -y --allow-unauthenticated upgrade
sudo apt-get -y --allow-unauthenticated autoremove
sudo apt-get -y --allow-unauthenticated autoclean

sudo service ntp stop
sudo ntpdate time.nist.gov
sudo service ntp start

echo "echo 1 > /proc/sys/vm/drop_caches" | sudo sh

##########################################################

##########################################################

sudo add-apt-repository ppa:yubico/stable
sudo apt-get -y --allow-unauthenticated install yubikey-neo-manager scdaemon

##########################################################

##########################################################

export CLOUD_SDK_REPO="cloud-sdk-$(lsb_release -c -s)"
echo "deb http://packages.cloud.google.com/apt $CLOUD_SDK_REPO main" | sudo tee /etc/apt/sources.list.d/google-cloud-sdk.list
curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
sudo apt-get update && sudo apt-get install google-cloud-sdk
gcloud init
gsutil config

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

GO_VERSION="1.6.2" && cd /usr/local && sudo rm -rf ./go && sudo curl -s https://storage.googleapis.com/golang/go$GO_VERSION.linux-amd64.tar.gz | sudo tar -v -C /usr/local/ -xz && cd $HOME;
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

cd $HOME && rm -rf $HOME/go-master
git clone https://go.googlesource.com/go $HOME/go-master
cd $HOME/go-master/src && ./make.bash
cd $HOME && $HOME/go-master/bin/go version

##########################################################

##########################################################

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
go get -v -u -f honnef.co/go/unused/cmd/unused && \
go get -v -u -f github.com/gyuho/psn && \
go get -v -u -f github.com/gyuho/gomp && \
go get -v -u -f github.com/coreos/dbtester && \
go get -v -u -f github.com/coreos/etcd-play && \
go get -v -u -f github.com/coreos/etcd

mkdir -p $HOME/go/src/github.com/gyuho
rm -rf $HOME/go/src/github.com/gyuho/learn
cd $HOME/go/src/github.com/gyuho
git clone git@github.com:gyuho/learn.git

mkdir -p $HOME/go/src/github.com/coreos
rm -rf $HOME/go/src/github.com/coreos/etcd
cd $HOME/go/src/github.com/coreos
git clone git@github.com:coreos/etcd.git

cd $GOPATH/src/github.com/nsf/gocode/vim && sudo ./update.sh

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

cd $HOME
rm -rf $HOME/ctags
git clone https://github.com/universal-ctags/ctags.git
cd ctags
./autogen.sh
./configure
make
sudo make install

go get -v -u -f github.com/jstemmer/gotags

##########################################################

##########################################################

wget https://github.com/google/protobuf/releases/download/v3.0.0-beta-2/protoc-3.0.0-beta-2-linux-x86_64.zip
unzip protoc-3.0.0-beta-2-linux-x86_64.zip
sudo cp ./protoc $GOPATH/protoc
rm -f /tmp/protoc.zip
protoc --version

##########################################################

##########################################################

# https://coreos.com/blog/getting-started-with-rkt-1.0.html
RKT_VERSION=1.0.0
rm -rf $HOME/rkt-v$RKT_VERSION
sudo curl -sf -o /tmp/rkt-v$RKT_VERSION.tar.gz -L https://github.com/coreos/rkt/releases/download/v$RKT_VERSION/rkt-v$RKT_VERSION.tar.gz
sudo tar -xzf /tmp/rkt-v$RKT_VERSION.tar.gz -C /tmp/
sudo mv /tmp/rkt-v$RKT_VERSION $HOME/rkt-v$RKT_VERSION

RKT_VERSION=1.4.0
rm -rf $HOME/rkt-v$RKT_VERSION
sudo curl -sf -o /tmp/rkt-v$RKT_VERSION.tar.gz -L https://github.com/coreos/rkt/releases/download/v$RKT_VERSION/rkt-v$RKT_VERSION.tar.gz
sudo tar -xzf /tmp/rkt-v$RKT_VERSION.tar.gz -C /tmp/
sudo mv /tmp/rkt-v$RKT_VERSION $HOME/rkt-v$RKT_VERSION

##########################################################

##########################################################

sudo apt-get -y update
sudo apt-get -y upgrade
sudo apt-get -y install apt-transport-https ca-certificates
sudo apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D

echo "deb https://apt.dockerproject.org/repo ubuntu-wily main" > a.temp
sudo mv a.temp /etc/apt/sources.list.d/docker.list

sudo apt-get -y update
sudo apt-get -y purge lxc-docker
sudo apt-cache policy docker-engine
sudo apt-get -y update
sudo apt-get -y install linux-image-extra-$(uname -r)

sudo apt-get -y install docker-engine
sudo service docker start
sudo docker version

sleep 3s
printf "\n\n"
sudo docker ps
sudo docker images
printf "\n\n"

##########################################################
