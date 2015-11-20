#!/bin/bash

# sudo fdisk -l
# /dev/sdb1
# sudo apt-get install pv
# dd if=archlinux-2015.11.01-dual.iso | pv | sudo dd of=/dev/sdb1

# reboot from USB

# https://wiki.archlinux.org/index.php/Installation_guide
# pacstrap /mnt base;
# genfstab -p /mnt >> /mnt/etc/fstab;
# arch-chroot /mnt;
# echo gyuho > /etc/hostname;
# ln -s /usr/share/zoneinfo/America/Los_Angeles /etc/localtime

# edit /etc/locale.gen
# local-gen

# localectl list-locales | grep en_;
# echo LANG=en_US.utf8 > /etc/locale.conf;

# mkinitcpio -p linux

# passwd



# Force pacman to refresh the package lists
pacman -Syyu;

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
