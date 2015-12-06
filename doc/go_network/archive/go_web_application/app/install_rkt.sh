#!/bin/bash

wget https://github.com/coreos/rkt/releases/download/v0.8.1/rkt-v0.8.1.tar.gz;
tar xzvf rkt-v0.8.1.tar.gz;
sudo cp rkt-v0.8.1/* .;

wget https://github.com/appc/spec/releases/download/v0.7.0/appc-v0.7.0.tar.gz;
tar xzvf appc-v0.7.0.tar.gz;
sudo cp appc-v0.7.0/actool .;
