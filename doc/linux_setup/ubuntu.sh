#!/bin/bash

# sudo visudo
# ubuntu ALL=(ALL) NOPASSWD: ALL

sudo apt-get -y install debconf-utils;
sudo apt-get -y install pcmanfm;
sudo apt-get -y install xclip;
sudo apt-get -y install lm-sensors;
sudo apt-get -y install tree;
sudo apt-get -y install htop;
sudo apt-get -y install ubuntu-restricted-extras update-manager-core;
sudo apt-get -y check;
sudo apt-get -y update;
sudo apt-get -y upgrade;
sudo apt-get -y autoclean;
sudo apt-get -y autoremove -f;
sudo sync && echo 3 | sudo tee /proc/sys/vm/drop_caches;
sudo ntpdate ntp.ubuntu.com;

sudo apt-get -y install git;
sudo apt-get -y install tmux;

sudo cp ./ubuntu_bashrc.sh $HOME/.bashrc && source $HOME/.bashrc;

sudo cp ./tmux.conf $HOME/.tmux.conf;
source $HOME/.tmux.conf;

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

sudo apt-get -y install postgresql;
sudo apt-get -y install mysql-server;
sudo echo mysql-server mysql-server/root_password password 1 | sudo debconf-set-selections;
sudo echo mysql-server mysql-server/root_password_again password 1 | sudo debconf-set-selections;
