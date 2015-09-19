#!/bin/bash
sudo docker build -t code .;
sudo docker run --publish 8080:8080 --name test --rm code;
