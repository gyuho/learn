
# install arch linux into USB disk
# https://wiki.archlinux.org/index.php/USB_flash_installation_media
# fdisk -l;
# /dev/sdbxY
sudo dd bs=4M if=$HOME/archlinux-2015.11.01-dual.iso of=/dev/sdx && sync;

# reboot from USB

# Boot Arch Linux x86_64

# wired
# ls /sys/class/net;
# ip link;
# systemctl enable dhcpcd@INTERFACENAME.service;
# INTERFACENAME is usually e*
# dhcpcd INTERFACENAME;
 
# wireless
# ls /sys/class/net;
# sudo pacman --noconfirm -S iw wpa_supplicant dialog wpa_actiond;
# ip link;
# systemctl enable dhcpcd@INTERFACENAME.service;
# INTERFACENAME is usually w*
# wifi-menu wlp3s0*;
# dhcpcd INTERFACENAME;

ping -c 3 www.google.com;

# check if it's efi
ls /sys/firmware/efi/efivars;
# if exists, it's efi



###########################
# INSTALL ONLY Arch Linux #
###########################
# https://wiki.archlinux.org/index.php/Beginners'_guide
# https://wiki.archlinux.org/index.php/Installation_guide

# find out what partitions you have
fdisk -l;
lsblk /dev/sdx;

# erase all
sgdisk --xap-all /dev/sdxY;

cfdisk;
# New -> Primary -> set Size (in MB) for Arch Linux
# Beginning, set Bootable
# New -> Primary -> set Size (in MB) for Linux Swap
# Write
# Quit

# format the first partition for Arch Linux
mkfs.ext4 /dev/sdxY;

# set the second partition for Linux Swap
mkswap /dev/sdxZ;
swapon /dev/sdxZ;

# check if the swap is on
lsblk /dev/sdx;

# mount the first partition
mount /dev/sdxY /mnt;

##################################

##################################
# INSTALL Arch Linux (DUAL-BOOT) #
##################################
# https://wiki.archlinux.org/index.php/Beginners'_guide
# https://wiki.archlinux.org/index.php/Installation_guide

# TO DO ...

##################################

# install basic libraries
pacstrap /mnt base base-devel;
pacman --noconfirm -Syyu;

# generate fstab to define how disk partitions
# should be mounted into the filesystem
genfstab -U /mnt > /mnt/etc/fstab;
vi /mnt/etc/fstab;

# chroot to /mnt, /bin/bash just for CLI
arch-chroot /mnt /bin/bash;

# install bootloader
pacman --noconfirm -Su grub grub-bios os-prober;
grub-install --recheck /dev/sdx;
grub-mkconfig -o /boot/grub/grub.cfg;

# set username
echo gyuho > /etc/hostname;
# vim /etc/hosts;

# set password
passwd

unmount /mnt;
exit;
reboot;

# set language
vim /etc/locale.gen; # remove # from en_US.UTF-8
locale-gen;
echo LANG=en_US.UTF-8 > /etc/locale.conf;
export LANG=en_US.UTF-8;

# set timezone
ln -s /usr/share/zoneinfo/America/Los_Angeles /etc/localtime;

# set hardware clock to UTC
hwclock --systohc --utc;
