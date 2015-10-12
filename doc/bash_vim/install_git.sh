#!/bin/bash

sudo apt-get -y install git;

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

