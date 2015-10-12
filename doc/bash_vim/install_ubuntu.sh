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

