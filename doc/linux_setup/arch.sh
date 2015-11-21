#!/bin/bash

# wifi-menu
pacman --noconfirm -Syyu;
timedatectl set-ntp true;

# ip link
pacman --noconfirm -S iw wpa_supplicant;
pacman --noconfirm -S dialog wpa_actiond;
# systemctl enable dhcpcd@INTERFACENAME.service;

# su
# nano /etc/sudoers
# gyuho ALL=(ALL) NOPASSWD: ALL
pacman -S gksu sudo;

x=$(lsb_release -a | grep "Distributor ID:")
if [ ${x:16:6} = "Ubuntu" ] || [ ${x:16:6} = "Debian" ]; then
    echo "Ubuntu ou Debian"
elif [ ${x:16:4} = "arch" ]; then
    echo "Arch linux"
else
    echo "Distro unknown!"
fi

sudo cp ./bashrc.sh $HOME/.bashrc && source $HOME/.bashrc;


pacman --noconfirm -S curl;
pacman --noconfirm -S wget;
pacman --noconfirm -S git;
pacman --noconfirm -S cmake;
pacman --noconfirm -S xterm;




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
