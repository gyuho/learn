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
# yaourt --noconfirm -S google-chrome;
# run with google-chrome-stable


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


printf "installing vim...\n\n" && sleep 5s;

sudo pacman --noconfirm -S clang;

sudo chown -R gyuho:gyuho $HOME/.vim;
sudo mkdir -p $HOME/.vim/bundle;
sudo mkdir -p $HOME/.vim/ftdetect;
sudo mkdir -p $HOME/.vim/syntax;
sudo chmod -R +x $HOME/.vim;
sudo git clone --progress \
	https://github.com/gmarik/Vundle.vim.git \
	~/.vim/bundle/Vundle.vim;

sudo cp ./vimrc.vim ~/.vimrc && \
source $HOME/.vimrc && \
sudo vim +PluginInstall +qall && \
sudo vim +PluginClean +qall;

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

sudo mkdir -p $HOME/.vim/bundle/YouCompleteMe && \
cd $HOME/.vim/bundle/YouCompleteMe && \
sudo ./install.sh --clang-completer --system-libclang;

cd $HOME && \
mkdir ycm_build && \
cd ycm_build;

sudo cmake -G "Unix Makefiles" . \
	~/.vim/bundle/YouCompleteMe/third_party/ycmd/cpp && \
sudo make ycm_support_libs;
