#!/bin/bash
printf "Installing tmux\n"
sudo apt-get -y install tmux;

printf "Copying tmux.conf\n"
sudo cp ./tmux.conf ~/.tmux.conf;
source ~/.tmux.conf;

