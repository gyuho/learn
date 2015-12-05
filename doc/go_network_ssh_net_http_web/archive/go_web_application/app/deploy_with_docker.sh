#!/bin/bash
sudo docker build -t app .;
sudo docker run --publish 8080:8080 --name test --rm app;
