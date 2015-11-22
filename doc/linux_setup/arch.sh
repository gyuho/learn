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
sudo pacman --noconfirm -S curl wget vim git feh;
sudo pacman --noconfirm -Syu && sudo pacman --noconfirm -S yaourt;
sudo mkdir -p $HOME/go/src/github.com/gyuho;
sudo mkdir -p $HOME/go/src/github.com/coreos;

# install default terminal
sudo pacman --noconfirm -S xfce4-terminal;

# install GUI
sudo pacman --noconfirm -S xorg xorg-xinit xorg-server xorg-twm xorg-xclock i3 dmenu && echo "exec i3" > $HOME/.xinitrc;
# sudo reboot;

sudo mkdir -p $HOME/Pictures && \
sudo wget https://wallpaperscraft.com/image/san_-_francisco_city_night_top_view_28432_1920x1200.jpg -q -O $HOME/Pictures/bg.jpg && \
exec --no-startup-id feh --bg-fill $HOME/Pictures/bg.jpg;

sudo mkdir -p $HOME/.i3;
sudo cp ./arch_pacman.conf /etc/pacman.conf;
sudo cp ./arch_xinitrc.conf $HOME/.xinitrc;
sudo cp ./arch_i3.conf $HOME/.i3/config;
sudo cp ./arch_i3status.conf $HOME/.i3/i3status.conf;
sudo cp ./arch_bashrc.sh $HOME/.bashrc;

sudo mkdir -p $HOME/Pictures && \
sudo wget https://wallpaperscraft.com/image/san_-_francisco_city_night_top_view_28432_1920x1200.jpg -q -O $HOME/Pictures/bg.jpg && \
feh --bg-scale $HOME/Pictures/bg.jpg;

# login
# startx;

# modkey + return      to start a terminal.
# modkey + shift + q   to close current window.
# modkey + f           switches the active window to fullscreen view.
# modkey + d           for dmenu.
# modkey + shift + e   to kill i3 session.

# install chrome
yaourt --noconfirm -S google-chrome;
# run with google-chrome-stable

