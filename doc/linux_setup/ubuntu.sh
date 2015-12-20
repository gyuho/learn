<<COMMENT
COMMENT

<<COMMENT
sudo su
nano /etc/sudoers
gyuho ALL=(ALL) NOPASSWD: ALL
COMMENT

sudo apt-get -y update && sudo apt-get -y upgrade && sudo apt-get -y autoremove && sudo apt-get -y autoclean;
sudo apt-get -y install build-essential;
sudo apt-get -y update && sudo apt-get -y upgrade && sudo apt-get -y autoremove && sudo apt-get -y autoclean;

<<COMMENT
- Time & Date
- Language Support
- https://www.google.com/chrome
- Software Updater

sudo reboot;

- Keyboard
COMMENT

#############################################################
printf "\n\n\n\n\ninstalling basics...\n\n" && sleep 1s;

sudo apt-get -y install terminator;
sudo apt-get -y install pcmanfm;
sudo apt-get -y install bash-completion;
sudo apt-get -y install git;
sudo apt-get -y install curl wget;
sudo apt-get -y install unzip gzip tar;
sudo apt-get -y install tree htop;
sudo apt-get -y install openssh;

sudo apt-get install -y python-pip python-dev build-essential;
sudo pip install --upgrade pip;

#############################################################
printf "\n\n\n\n\ninstalling git...\n\n" && sleep 1s;

echo "[user]
  email = gyuhox@gmail.com
  name = Gyu-Ho Lee

[color]
  diff = auto
  status = auto
  branch = auto
  ui = auto" > $HOME/.gitconfig;

git config --global user.name "Gyu-Ho Lee";
git config --global user.email "gyuhox@gmail.com";
git config --global core.editor "vim";

#############################################################
printf "\n\n\n\n\ninstalling protobuf...\n\n" && sleep 1s;

cd $HOME/go/src/github.com && rm -rf google/protobuf && mkdir google;
cd $HOME/go/src/github.com/google && git clone https://github.com/google/protobuf.git;
cd $HOME/go/src/github.com/google/protobuf && ./autogen.sh;
cd $HOME/go/src/github.com/google/protobuf && ./configure;
cd $HOME/go/src/github.com/google/protobuf && make;
cd $HOME/go/src/github.com/google/protobuf && make check;
cd $HOME/go/src/github.com/google/protobuf && make install;

#############################################################
printf "\n\n\n\n\ninstalling go...\n\n" && sleep 1s;

cd $HOME && \
mkdir -p $HOME/go/src && \
mkdir -p $HOME/go/src/github.com && \
mkdir -p $HOME/go/src/github.com/coreos && \
mkdir -p $HOME/go/src/github.com/gyuho && \
mkdir -p $HOME/go/src/golang.org;

cd $HOME && sudo rm -rf go1.4;
cd $HOME && sudo rm -rf go1.4_temp && mkdir -p $HOME/go1.4_temp;
sudo curl -s https://storage.googleapis.com/golang/go1.4.linux-amd64.tar.gz | sudo tar -v -C $HOME/go1.4_temp -xz;
cd $HOME/go1.4_temp && sudo mv ./go ./go1.4 && sudo mv ./go1.4 ..;
cd $HOME && sudo rm -rf $HOME/go1.4_temp;

cd $HOME && rm -rf ./go-master && mkdir -p $HOME/go-master;
cd $HOME/go-master && git clone https://go.googlesource.com/go && cd $HOME/go-master/go/src && ./all.bash;

cd /usr/local && sudo rm -rf ./go;
sudo curl -s https://storage.googleapis.com/golang/go1.5.2.linux-amd64.tar.gz | sudo tar -v -C /usr/local/ -xz;

if grep -q GOPATH "$(echo $HOME)/.bashrc"; then 
	echo "bashrc already has GOPATH...";
else
	echo "adding GOPATH to bashrc...";
	echo "export GOPATH=$(echo $HOME)/go" >> $HOME/.bashrc && \
	PATH_VAR=$PATH":/usr/local/go/bin:$(echo $HOME)/go/bin" && \
	echo "export PATH=$(echo $PATH_VAR)" >> $HOME/.bashrc && \
	source $HOME/.bashrc;
fi

cd $HOME && \
printf "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Successfully installed Go.\")\n}" > $HOME/temp.go;
cd $HOME && go run temp.go && rm -f temp.go && go version;

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
go get -v -u github.com/alecthomas/gometalinter && \
cd $GOPATH/src/github.com/nsf/gocode/vim && sudo ./update.sh && \
cd $HOME;

#############################################################
printf "\n\n\n\n\ninstalling vim...\n\n" && sleep 1s;

sudo apt-get -y install gvim vim;

sudo chown -R gyuho:gyuho $HOME/.vim;
sudo mkdir -p $HOME/.vim/bundle;
sudo mkdir -p $HOME/.vim/ftdetect;
sudo mkdir -p $HOME/.vim/syntax;
sudo chmod -R +x $HOME/.vim;
sudo git clone --progress https://github.com/gmarik/Vundle.vim.git ~/.vim/bundle/Vundle.vim;

cd $HOME/go/src/github.com/gyuho/learn/doc/linux_setup;
sudo cp ./vimrc.vim ~/.vimrc;
source $HOME/.vimrc;
vim +PluginInstall +qall;
vim +PluginClean +qall;

# :GoInstallBinaries

#############################################################
printf "\n\n\n\n\ndone!\n\n\n\n\n"

# free pagecache, dentries, inodes / 1 for pagecache
# and the system starts caching immediately again
sudo sync && echo 3 | sudo tee /proc/sys/vm/drop_caches;
