#!/bin/bash
sudo apt-get -y install postgresql;
sudo apt-get -y install mysql-server;
sudo echo mysql-server mysql-server/root_password password 1 | sudo debconf-set-selections;
sudo echo mysql-server mysql-server/root_password_again password 1 | sudo debconf-set-selections;

