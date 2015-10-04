#!/bin/bash
printf "Installing vim, tmux\n"
sudo apt-get -y install vim vim-gnome tmux;

printf "Copying tmux.conf\n"
sudo cp ./tmux.conf ~/.tmux.conf;
source ~/.tmux.conf;

printf "Creating directories in .vim\n"
sudo mkdir -p $HOME/.vim/bundle;
sudo mkdir -p $HOME/.vim/ftdetect;
sudo mkdir -p $HOME/.vim/syntax;
sudo chmod -R +x $HOME/.vim;
sudo git clone --progress \
	https://github.com/gmarik/Vundle.vim.git \
	~/.vim/bundle/Vundle.vim;

printf "Copying vimrc\n"
sudo cp ./vimrc.vim ~/.vimrc;
source $HOME/.vimrc;
sudo vim +PluginInstall +qall;
sudo vim +PluginClean +qall;

# https://github.com/Valloric/YouCompleteMe
sudo apt-get install cmake;
sudo apt-get install python-dev;

# cd $HOME/.vim/bundle/YouCompleteMe;
# sudo ./install.sh --clang-completer --system-libclang;

cd $HOME;
mkdir ycm_build;
cd ycm_build;

sudo cmake -G "Unix Makefiles" . \
	~/.vim/bundle/YouCompleteMe/third_party/ycmd/cpp;
sudo make ycm_support_libs;

