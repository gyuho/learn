#!/bin/bash

# sudo fdisk -l
# /dev/sdb1
# sudo apt-get install pv
# dd if=archlinux-2015.11.01-dual.iso | pv | sudo dd of=/dev/sdb1
# reboot from USB

sudo cp ./bashrc.sh $HOME/.bashrc && source $HOME/.bashrc;

# su
# nano /etc/sudoers
# gyuho ALL=(ALL) NOPASSWD: ALL
pacman -S gksu sudo;

pacman --noconfirm -S bash;
pacman --noconfirm -S curl;
pacman --noconfirm -S git;
pacman --noconfirm -S cmake;
pacman --noconfirm -S grep;
