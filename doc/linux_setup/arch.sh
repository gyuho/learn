# su
# echo "blacklist pcspkr" > /etc/modprobe.d/nobeep.conf;

mkdir -p $HOME/go/src/github.com/gyuho;
mkdir -p $HOME/go/src/github.com/coreos;

# wired
# ip link;
# systemctl enable dhcpcd@INTERFACENAME.service;
# INTERFACENAME is usually e*
# dhcpcd INTERFACENAME;

# wireless
# ls /sys/class/net;

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

#############################################################
printf "\n\n\n\n\ninstalling basics...\n\n" && sleep 1s;

sudo pacman --noconfirm -Su sudo;
sudo pacman --noconfirm -Su bash-completion;
sudo pacman --noconfirm -Su git;
sudo pacman --noconfirm -Su curl wget;
sudo pacman --noconfirm -Su gvim vim;
sudo pacman --noconfirm -Su unzip gzip tar;
sudo pacman --noconfirm -Su dbus;
sudo pacman --noconfirm -Su tree htop;
sudo pacman --noconfirm -Su openssh;
sudo pacman --noconfirm -Su netctl;
sudo pacman --noconfirm -Su iw;
sudo pacman --noconfirm -Su dialog;
sudo pacman --noconfirm -Su wpa_actiond;
sudo pacman --noconfirm -Su wpa_supplicant;
sudo pacman --noconfirm -Su systemd;
sudo pacman --noconfirm -Su systemd-arch-units;
sudo pacman --noconfirm -Su networkmanager;
sudo pacman --noconfirm -Su net-tools;
sudo pacman --noconfirm -Su gnu-netcat;
sudo pacman --noconfirm -Su ntp;
sudo pacman --noconfirm -Su python-pip;

timedatectl set-ntp true;
sudo systemctl enable ntpd;
sudo /etc/rc.d/ntpd start;

sudo pacman --noconfirm -Syu && sudo pacman --noconfirm -Su yaourt;
sudo pacman --noconfirm -Syu && sudo pacman --noconfirm -Rns $(sudo pacman -Qtdq);

#############################################################
printf "\n\n\n\n\ninstalling gui...\n\n" && sleep 1s;

sudo pacman --noconfirm -Su pcmanfm;
sudo pacman --noconfirm -Su zathura;
sudo pacman --noconfirm -Su zathura-pdf-poppler;
sudo pacman --noconfirm -Su zathura-djvu;

sudo pacman --noconfirm -Su terminator;
sudo pacman --noconfirm -Su transmission-daemon;
sudo pacman --noconfirm -Su transmission-qt;
sudo pacman --noconfirm -Su gnome-screenshot;
sudo pacman --noconfirm -Su vlc;

<<COMMENT

yaourt --noconfirm -Sb ttf-nanum;
sudo pacman --noconfirm -Su noto-fonts;
sudo pacman --noconfirm -Su noto-fonts-cjk;
sudo pacman --noconfirm -Su noto-fonts-emoji;
sudo pacman --noconfirm -Su ttf-dejavu;
sudo pacman --noconfirm -Su ttf-droid;
sudo pacman --noconfirm -Su libhangul;
sudo pacman --noconfirm -Su ttf-baekmuk;
sudo pacman --noconfirm -Su libreoffice-fresh;
sudo pacman --noconfirm -Su hunspell-en;
sudo pacman --noconfirm -Su hunspell;

COMMENT

sudo pacman --noconfirm -Su xorg;
sudo pacman --noconfirm -Su xorg-xinit;
sudo pacman --noconfirm -Su xorg-server;
sudo pacman --noconfirm -Su xorg-utils;
sudo pacman --noconfirm -Su xorg-twm;
sudo pacman --noconfirm -Su xorg-xclock;
sudo pacman --noconfirm -Su lxde;
sudo pacman --noconfirm -Su lxdm;
sudo pacman --noconfirm -Su alsa-utils;
sudo pacman --noconfirm -Syu && sudo pacman --noconfirm -Rns $(sudo pacman -Qtdq);

# sudo reboot;

sudo mkdir -p $HOME/fontconfig;
cd $HOME/go/src/github.com/gyuho/learn/doc/linux_setup;
sudo cp ./arch_pacman.conf /etc/pacman.conf;
sudo cp ./arch_xinitrc.conf $HOME/.xinitrc && sudo chmod +x $HOME/.xinitrc;
sudo cp ./bashrc.sh $HOME/.bashrc && source $HOME/.bashrc;
sudo cp ./arch_fonts.conf $HOME/fontconfig/fonts.conf;
sudo cp ./arch_etc_rc.conf /etc/rc.conf;
sudo cp ./arch_asoundrc.conf $HOME/.asoundrc;
sudo cp ./arch_terminator.conf $HOME/.config/terminator/config;
sudo cp ./arch_lxde_shortcuts.xml $HOME/.config/openbox/lxde-rc.xml;

# login
# startx;

#############################################################
printf "\n\n\n\n\ninstalling chromium...\n\n" && sleep 1s;

# install chrome
sudo pacman --noconfirm -Su chromium;

# yaourt --noconfirm -S google-chrome;
# google-chrome-stable

#############################################################
printf "\n\n\n\n\ninstalling protobuf...\n\n" && sleep 1s;

cd $HOME/go/src/github.com && rm -rf google/protobuf && mkdir google;
cd $HOME/go/src/github.com/google && git clone https://github.com/google/protobuf.git;
cd $HOME/go/src/github.com/google/protobuf && ./autogen.sh;
cd $HOME/go/src/github.com/google/protobuf && ./configure;
cd $HOME/go/src/github.com/google/protobuf && make;
cd $HOME/go/src/github.com/google/protobuf && make check;
cd $HOME/go/src/github.com/google/protobuf && make install;

#############################################################
printf "\n\n\n\n\ninstalling git...\n\n" && sleep 1s;

echo "[user]
  email = gyuhox@gmail.com
  name = Gyu-Ho Lee

[color]
  diff = auto
  status = auto
  branch = auto
  ui = auto" > $HOME/.gitconfig;

git config --global user.name "Gyu-Ho Lee";
git config --global user.email "gyuhox@gmail.com";
git config --global core.editor "vim";

#############################################################
printf "\n\n\n\n\ninstalling vim...\n\n" && sleep 1s;

sudo pacman --noconfirm -Su clang;

sudo chown -R gyuho:gyuho $HOME/.vim;
sudo mkdir -p $HOME/.vim/bundle;
sudo mkdir -p $HOME/.vim/ftdetect;
sudo mkdir -p $HOME/.vim/syntax;
sudo chmod -R +x $HOME/.vim;
sudo git clone --progress https://github.com/gmarik/Vundle.vim.git ~/.vim/bundle/Vundle.vim;

cd $HOME;

cd $HOME/go/src/github.com/gyuho/learn/doc/linux_setup;
sudo cp ./vimrc.vim ~/.vimrc;
source $HOME/.vimrc;
vim +PluginInstall +qall;
vim +PluginClean +qall;
# :GoInstallBinaries

#############################################################
printf "\n\n\n\n\ninstalling go...\n\n" && sleep 1s;

cd $HOME && \
mkdir -p $HOME/go/src && \
mkdir -p $HOME/go/src/github.com && \
mkdir -p $HOME/go/src/github.com/coreos && \
mkdir -p $HOME/go/src/github.com/gyuho && \
mkdir -p $HOME/go/src/golang.org;

cd $HOME && sudo rm -rf go1.4 && \
cd $HOME && sudo rm -rf go1.4_temp && mkdir -p $HOME/go1.4_temp && \
sudo curl -s https://storage.googleapis.com/golang/go1.4.linux-amd64.tar.gz | sudo tar -v -C $HOME/go1.4_temp -xz && \
cd $HOME/go1.4_temp && sudo mv ./go ./go1.4 && sudo mv ./go1.4 .. && \
cd $HOME && sudo rm -rf $HOME/go1.4_temp;

cd $HOME && rm -rf ./go-master && mkdir -p $HOME/go-master && \
cd $HOME/go-master && git clone https://go.googlesource.com/go && cd $HOME/go-master/go/src && ./all.bash;

cd /usr/local && sudo rm -rf ./go;
sudo curl -s https://storage.googleapis.com/golang/go1.5.2.linux-amd64.tar.gz | sudo tar -v -C /usr/local/ -xz;

if grep -q GOPATH "$(echo $HOME)/.bashrc"; then 
	echo "bashrc already has GOPATH...";
else
	echo "adding GOPATH to bashrc...";
	echo "export GOPATH=$(echo $HOME)/go" >> $HOME/.bashrc && \
	PATH_VAR=$PATH":/usr/local/go/bin:$(echo $HOME)/go/bin" && \
	echo "export PATH=$(echo $PATH_VAR)" >> $HOME/.bashrc && \
	source $HOME/.bashrc;
fi

cd $HOME && \
printf "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Successfully installed Go.\")\n}" > $HOME/temp.go;
cd $HOME && go run temp.go && rm -f temp.go && go version;

cd $HOME && \
go get -v -u github.com/tools/godep && \
go get -v -u golang.org/x/tools/cmd/... && \
go get -v -u github.com/golang/lint/golint && \
go get -v -u github.com/nsf/gocode && \
go get -v -u github.com/motain/gocheck && \
go get -v -u github.com/vaughan0/go-ini && \
go get -v -u github.com/rogpeppe/godef && \
go get -v -u github.com/kisielk/errcheck && \
go get -v -u github.com/jstemmer/gotags && \
go get -v -u github.com/alecthomas/gometalinter && \
cd $GOPATH/src/github.com/nsf/gocode/vim && sudo ./update.sh && \
cd $HOME;

#############################################################
printf "\n\n\n\n\nDONE\n\n\n\n\n"

sudo pacman --noconfirm -Syu && sudo pacman --noconfirm -Rns $(sudo pacman -Qtdq);

# free pagecache, dentries, inodes / 1 for pagecache
# and the system starts caching immediately again
sudo sync && echo 3 | sudo tee /proc/sys/vm/drop_caches;

# sudo echo mysql-server mysql-server/root_password password 1 | sudo debconf-set-selections;
# sudo echo mysql-server mysql-server/root_password_again password 1 | sudo debconf-set-selections;
