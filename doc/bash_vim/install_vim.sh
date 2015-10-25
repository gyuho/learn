#!/bin/bash

# printf "Installing vim\n"
# sudo apt-get -y build-dep vim;
# sudo apt-get -y install ncurses-dev;
# sudo apt-get -y install vim vim-gnome;
# cd $HOME && rm -rf vim && git clone https://github.com/vim/vim.git;
# cd $HOME && cd vim/src && sudo make;

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

sudo apt-get -y install ctags;
cd $HOME/go && ctags -R ./*;
cd $HOME;

# https://github.com/Valloric/YouCompleteMe
sudo apt-get -y install cmake;
sudo apt-get -y install python-dev;

# cd $HOME/.vim/bundle/YouCompleteMe;
# sudo ./install.sh --clang-completer --system-libclang;

cd $HOME;
mkdir ycm_build;
cd ycm_build;

sudo cmake -G "Unix Makefiles" . \
	~/.vim/bundle/YouCompleteMe/third_party/ycmd/cpp;
sudo make ycm_support_libs;

