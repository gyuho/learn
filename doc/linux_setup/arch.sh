timedatectl set-ntp true;

# wired
# ip link;
# systemctl enable dhcpcd@INTERFACENAME.service;
# INTERFACENAME is usually e*
# dhcpcd INTERFACENAME;
# 
# wireless
# ls /sys/class/net;
sudo pacman --noconfirm -S iw wpa_supplicant dialog wpa_actiond;
# ip link;
# systemctl enable dhcpcd@INTERFACENAME.service;
# INTERFACENAME is usually w*
# wifi-menu wlp3s0*;
# dhcpcd INTERFACENAME;

# useradd -m gyuho;
# passwd gyuho;
# reboot;


# su
# nano /etc/sudoers
# gyuho ALL=(ALL) NOPASSWD: ALL
sudo pacman --noconfirm -S sudo;
sudo pacman --noconfirm -Syu && sudo pacman --noconfirm -Rns $(sudo pacman -Qtdq);
sudo pacman --noconfirm -S curl wget vim git xterm;

# install GUI
sudo pacman --noconfirm -S xorg xorg-xinit xorg-server i3 dmenu && echo "exec i3" > $HOME/.xinitrc;
# sudo reboot;

# login
# startx;

# sudo vi /etc/pacman.conf;
echo '
[archlinuxfr]
SigLevel = Never
Server = http://repo.archlinux.fr/$arch
' >> /etc/pacman.conf;

mkdir -p $HOME/go/src/github.com/gyuho;
mkdir -p $HOME/go/src/github.com/coreos;
