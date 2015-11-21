#!/bin/bash

timedatectl set-ntp true;

# wired
# ip link;
# systemctl enable dhcpcd@INTERFACENAME.service;
# INTERFACENAME is usually e*
# dhcpcd INTERFACENAME;
# 
# wireless
# ls /sys/class/net;
pacman --noconfirm -S iw wpa_supplicant dialog wpa_actiond;
# ip link;
# systemctl enable dhcpcd@INTERFACENAME.service;
# INTERFACENAME is usually w*
# wifi-menu wlp3s0*;
# dhcpcd INTERFACENAME;

useradd -m gyuho;
passwd gyuho;
reboot;




# su
# nano /etc/sudoers
# gyuho ALL=(ALL) NOPASSWD: ALL
sudo pacman --noconfirm -S sudo;
sudo pacman --noconfirm -Syu && sudo pacman --noconfirm -Rns $(sudo pacman -Qtdq);
sudo pacman --noconfirm -S curl wget vim git;

# install GUI
sudo pacman --noconfirm -S xorg xorg-xinit xorg-server i3 dmenu && echo "exec i3" > $HOME/.xinitrc;
sudo reboot;

# login
startx;

# install Chrome















pacman --noconfirm -S cmake;
pacman --noconfirm -S xterm;





x=$(lsb_release -a | grep "Distributor ID:")
if [ ${x:16:6} = "Ubuntu" ] || [ ${x:16:6} = "Debian" ]; then
    echo "Ubuntu ou Debian"
elif [ ${x:16:4} = "arch" ]; then
    echo "Arch linux"
else
    echo "Distro unknown!"
fi

sudo cp ./bashrc.sh $HOME/.bashrc && source $HOME/.bashrc;





# ============================================================
# Get the URL from https://aur.archlinux.org/packages/libgcrypt15/
# ============================================================
cd /usr/src
sudo wget https://aur.archlinux.org/packages/li/libgcrypt15/libgcrypt15.tar.gz
sudo tar zxf libgcrypt15.tar.gz
cd libgcrypt15
sudo makepkg -s --asroot
# >>>>> be very careful about next step because this may be different <<<<<
sudo pacman -U libgcrypt*

# ============================================================
# Get the URL from https://aur.archlinux.org/packages/google-chrome/
# ============================================================
cd /usr/src/
sudo wget https://aur.archlinux.org/packages/go/google-chrome/google-chrome.tar.gz
sudo tar zxf google-chrome.tar.gz
cd google-chrome
sudo makepkg -s --asroot
# >>>>> be very careful about next step because this may be different <<<<<
sudo pacman -U google-chrome*
