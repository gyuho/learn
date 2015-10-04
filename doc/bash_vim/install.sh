#!/bin/bash

# sudo visudo
# ubuntu ALL=(ALL) NOPASSWD: ALL

sudo apt-get -y install vim vim-gnome tmux;
sudo apt-get -y install debconf-utils;

mkdir -p $HOME/go/src/github.com;
mkdir -p $HOME/go/src/golang.org;
sudo mkdir -p $HOME/go/src/github.com/gyuho;

sudo apt-get -y install postgresql;
sudo apt-get -y install mysql-server;
sudo echo mysql-server mysql-server/root_password password 1 | sudo debconf-set-selections;
sudo echo mysql-server mysql-server/root_password_again password 1 | sudo debconf-set-selections;

sudo apt-get -y install pcmanfm;
sudo apt-get -y install xclip;
sudo apt-get -y install git;
sudo apt-get -y install lm-sensors;
sudo apt-get -y install tree;

# git
ssh-keygen -t rsa -C "gyuhox@gmail.com" -f $HOME/.ssh/id_rsa -N "";
eval "$(ssh-agent -s)";
ssh-add /home/ubuntu/.ssh/id_rsa;
xclip -sel clip < $HOME/.ssh/id_rsa.pub;

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

sudo apt-get -y install ubuntu-restricted-extras update-manager-core;
sudo apt-get -y check;
sudo apt-get -y update;
sudo apt-get -y upgrade;
sudo apt-get -y autoclean;
sudo apt-get -y autoremove -f;
sudo sync && echo 3 | sudo tee /proc/sys/vm/drop_caches;
sudo ntpdate ntp.ubuntu.com;

# vim
cd $HOME;
sudo mkdir -p $HOME/.vim/bundle;
sudo mkdir -p $HOME/.vim/ftdetect;
sudo mkdir -p $HOME/.vim/syntax;
sudo chmod -R +x $HOME/.vim;
sudo git clone --progress https://github.com/gmarik/Vundle.vim.git ~/.vim/bundle/Vundle.vim;


