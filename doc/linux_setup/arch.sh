#!/bin/bash

# sudo fdisk -l
# /dev/sdb1
# sudo apt-get install pv
# dd if=archlinux-2015.11.01-dual.iso | pv | sudo dd of=/dev/sdb1
# reboot from USB

x=$(lsb_release -a | grep "Distributor ID:")
if [ ${x:16:6} = "Ubuntu" ] || [ ${x:16:6} = "Debian" ]; then
    echo "Ubuntu ou Debian"
elif [ ${x:16:4} = "arch" ]; then
    echo "Arch linux"
else
    echo "Distro unknown!"
fi

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
