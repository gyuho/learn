sudo pacman --noconfirm -Syu && sudo pacman --noconfirm -Rns $(sudo pacman -Qtdq);
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

#############################################################
printf "\n\n\n\n\ninstalling basics...\n\n" && sleep 5s;

sudo pacman --noconfirm -S sudo;
sudo pacman --noconfirm -Syu && sudo pacman --noconfirm -Rns $(sudo pacman -Qtdq);
sudo pacman --noconfirm -S curl wget gvim vim git;
sudo pacman --noconfirm -Syu && sudo pacman --noconfirm -S yaourt;

sudo mkdir -p $HOME/go/src/github.com/gyuho;
sudo mkdir -p $HOME/go/src/github.com/coreos;

sudo pacman --noconfirm -S alsa-utils;
sudo pacman --noconfirm -S dbus;

sudo pacman --noconfirm -S terminator xterm;
sudo pacman --noconfirm -S pcmanfm;

sudo pacman --noconfirm -S noto-fonts noto-fonts-cjk noto-fonts-emoji;
sudo pacman --noconfirm -S ttf-dejavu;
sudo pacman --noconfirm -S ttf-droid;
sudo pacman --noconfirm -S ttf-baekmuk;

# terminator
# Ctrl-Shift-E: will split the view vertically.
# Ctrl-Shift-O: will split the view horizontally.
# Ctrl-Shift-P: will focus be active on the previous view.
# Ctrl-Shift-N: will focus be active on the next view.
# Ctrl-Shift-W: will close the view where the focus is on.
# Ctrl-Shift-Q: will exit terminator.
# Ctrl-Shift-X: will focus active window and  enlarge it

#############################################################
printf "\n\n\n\n\ninstalling gui...\n\n" && sleep 5s;

sudo pacman --noconfirm -S xorg xorg-xinit xorg-server \
	xorg-utils xorg-twm xorg-xclock \
	xfce4 xfce4-goodies;

sudo pacman --noconfirm -Syu && sudo pacman --noconfirm -Rns $(sudo pacman -Qtdq);

# sudo reboot;

sudo mkdir -p $HOME/fontconfig;
sudo cp ./arch_pacman.conf /etc/pacman.conf;
sudo cp ./arch_xinitrc.conf $HOME/.xinitrc && sudo chmod +x $HOME/.xinitrc;
sudo cp ./arch_bashrc.sh $HOME/.bashrc;
sudo cp ./arch_fonts.conf $HOME/fontconfig/fonts.conf;

# login
# startx;

#############################################################
printf "\n\n\n\n\ninstalling chrome...\n\n" && sleep 5s;

# install chrome
yaourt --noconfirm -S google-chrome;
# run with google-chrome-stable

#############################################################
printf "\n\n\n\n\ninstalling git...\n\n" && sleep 5s;

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
printf "\n\n\n\n\ninstalling vim...\n\n" && sleep 5s;

sudo pacman --noconfirm -S clang;

sudo chown -R gyuho:gyuho $HOME/.vim;
sudo mkdir -p $HOME/.vim/bundle;
sudo mkdir -p $HOME/.vim/ftdetect;
sudo mkdir -p $HOME/.vim/syntax;
sudo chmod -R +x $HOME/.vim;
sudo git clone --progress https://github.com/gmarik/Vundle.vim.git ~/.vim/bundle/Vundle.vim;

sudo cp ./vimrc.vim ~/.vimrc && \
source $HOME/.vimrc && \
vim +PluginInstall +qall && \
vim +PluginClean +qall;

sudo pacman --noconfirm -S ctags && \
cd $HOME/go && ctags -R ./* && \
cd $HOME;

sudo mkdir -p $HOME/.vim/ctags && \
cd $HOME/.vim/ctags && \
pacman -Ql glibc | awk '/\/usr\/include/{print $2}' > c_headers && \
ctags -L c_headers --c-kinds=+p --fields=+iaS --extra=+q -f c && \
pacman -Ql gcc | awk '/\/usr\/include/{print $2}' > c++_headers && \
ctags -L c++_headers --c++-kinds=+p --fields=+iaS --extra=+q -f c++;

# https://github.com/Valloric/YouCompleteMe
sudo pacman --noconfirm -S cmake && \
sudo pacman --noconfirm -S python;

# sudo mkdir -p $HOME/.vim/bundle/YouCompleteMe && \
# cd $HOME/.vim/bundle/YouCompleteMe && \
# sudo ./install.sh --clang-completer --system-libclang;

cd $HOME && \
mkdir ycm_build && \
cd ycm_build;

sudo cmake -G "Unix Makefiles" . \
	~/.vim/bundle/YouCompleteMe/third_party/ycmd/cpp && \
sudo make ycm_support_libs;

#############################################################
printf "\n\n\n\n\ninstalling go...\n\n" && sleep 5s;

cd $HOME && \
mkdir -p $HOME/go/src && \
mkdir -p $HOME/go/src/github.com && \
mkdir -p $HOME/go/src/github.com/coreos && \
mkdir -p $HOME/go/src/github.com/gyuho && \
mkdir -p $HOME/go/src/golang.org;

cd /usr/local && sudo rm -rf ./go && \
sudo curl -s https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz | sudo tar -v -C /usr/local/ -xz;

echo "export GOPATH=$(echo $HOME)/go" >> $HOME/.bashrc && \
PATH_VAR=$PATH":/usr/local/go/bin:$(echo $HOME)/go/bin" && \
echo "export PATH=$(echo $PATH_VAR)" >> $HOME/.bashrc && \
source $HOME/.bashrc;

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
cd $GOPATH/src/github.com/nsf/gocode/vim && sudo ./update.sh && \
cd $HOME;

#############################################################
printf "\n\n\n\n\nDONE\n\n\n\n\n"
sudo pacman --noconfirm -Syu && sudo pacman --noconfirm -Rns $(sudo pacman -Qtdq);

