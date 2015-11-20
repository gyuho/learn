#!/bin/bash

# https://www.python.org/downloads/

cd $HOME;
sudo apt-get -y install python-pip python-dev python-all \
python-psycopg2 python-numpy python-pandas python-mysqldb;

sudo pip install --upgrade pip;
sudo pip install --upgrade psycopg2;
sudo pip install --upgrade pyyaml;
sudo pip install --upgrade gevent;
sudo pip install --upgrade sqlalchemy;
sudo pip install --upgrade boto;