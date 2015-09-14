[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# bash, vim

I primarily work with Ubuntu, bash, vim.

- [Reference](#reference)
- [install](#install)
- [ubuntu](#ubuntu)
- [terminal](#terminal)
- [tmux](#tmux)
- [`vim`](#vim)

[↑ top](#bash-vim)
<br><br><br><br>
<hr>





#### Reference

- [The art of command line](https://github.com/jlevy/the-art-of-command-line)

[↑ top](#bash-vim)
<br><br><br><br>
<hr>






#### install

```sh
#!/bin/bash

# sudo visudo
# ubuntu ALL=(ALL) NOPASSWD: ALL

sudo apt-get -y install vim vim-gnome tmux;
sudo apt-get -y install debconf-utils;

mkdir -p $HOME/go/src/github.com;
mkdir -p $HOME/go/src/golang.org;
sudo mkdir -p $HOME/go/src/github.com/gyuho;

sudo apt-get -y install postgresql;
sudo apt-get -y install mysql-server;
sudo echo mysql-server mysql-server/root_password password 1 | sudo debconf-set-selections;
sudo echo mysql-server mysql-server/root_password_again password 1 | sudo debconf-set-selections;

sudo apt-get -y install pcmanfm;
sudo apt-get -y install xclip;
sudo apt-get -y install git;
sudo apt-get -y install lm-sensors;

# git
ssh-keygen -t rsa -C "gyuhox@gmail.com" -f $HOME/.ssh/id_rsa -N "";
eval "$(ssh-agent -s)";
ssh-add /home/ubuntu/.ssh/id_rsa;
xclip -sel clip < $HOME/.ssh/id_rsa.pub;

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

sudo apt-get -y install ubuntu-restricted-extras update-manager-core;
sudo apt-get -y check;
sudo apt-get -y update;
sudo apt-get -y upgrade;
sudo apt-get -y autoclean;
sudo apt-get -y autoremove -f;
sudo sync && echo 3 | sudo tee /proc/sys/vm/drop_caches;
sudo ntpdate ntp.ubuntu.com;

# vim
cd $HOME;
sudo mkdir -p $HOME/.vim/bundle;
sudo mkdir -p $HOME/.vim/ftdetect;
sudo mkdir -p $HOME/.vim/syntax;
sudo chmod -R +x $HOME/.vim;
sudo git clone --progress https://github.com/gmarik/Vundle.vim.git ~/.vim/bundle/Vundle.vim;


# go
cd $HOME;
mkdir -p $HOME/go/src/github.com;
mkdir -p $HOME/go/src/golang.org;
sudo curl -s https://storage.googleapis.com/golang/go1.5.linux-amd64.tar.gz | sudo tar -v -C /usr/local/ -xz;

echo "export GOPATH=$(echo $HOME)/go" >> $HOME/.bashrc;
TEMP_PATH=$PATH':/usr/local/go/bin:/home/ubuntu/go/bin'
echo "export PATH=$(echo $TEMP_PATH)" >> $HOME/.bashrc;
source $HOME/.bashrc;

cd $HOME;
printf "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Successfully installed Go\")\n}" > $HOME/temp.go; 
go run temp.go; 
rm -rf temp.go;
go version;

go get -v github.com/tools/godep;
go get -v github.com/lib/pq;
go get -v github.com/go-sql-driver/mysql;
go get -v golang.org/x/tools/cmd/goimports;
go get -v github.com/golang/lint/golint;
go get -v github.com/nsf/gocode;
go get -v github.com/motain/gocheck;
go get -v github.com/vaughan0/go-ini;
cd $GOPATH/src/github.com/nsf/gocode/vim; sudo ./update.sh;


# python
sudo apt-get -y install python-pip python-dev python-all \
python-psycopg2 python-numpy python-pandas python-mysqldb;

sudo pip install --upgrade pip;
sudo pip install --upgrade psycopg2;
sudo pip install --upgrade pyyaml;
sudo pip install --upgrade gevent;
sudo pip install --upgrade sqlalchemy;
sudo pip install --upgrade boto;
```

[↑ top](#bash-vim)
<br><br><br><br>
<hr>











#### ubuntu

- <kbd>ctrl</kbd> + <kbd>alt</kbd> + <kbd>t</kbd> : `launch terminal`
- <kbd>print</kbd> : `Take a screenshot of an area`
- <kbd>alt</kbd> + <kbd>space</kbd> : `switch to previous (input) source`
- <kbd>ctrl</kbd> + <kbd>alt</kbd> + <kbd>m</kbd> : `maximize window`
- <kbd>ctrl</kbd> + <kbd>alt</kbd> + <kbd>r</kbd> : `restore window`
- <kbd>alt</kbd> + <kbd>q</kbd> : `close window`
- <kbd>ctrl</kbd> + <kbd>alt</kbd> + <kbd>v</kbd> : `maximize window vertically`

[↑ top](#bash-vim)
<br><br><br><br>
<hr>







#### terminal

- <kbd>ctrl</kbd> + <kbd>t</kbd> : `new tab`
- <kbd>ctrl</kbd> + <kbd>w</kbd> : `close tab`
- <kbd>alt</kbd> + <kbd>q</kbd> : `close window`
- <kbd>ctrl</kbd> + <kbd>c</kbd> : `copy`
- <kbd>ctrl</kbd> + <kbd>v</kbd> : `paste`
- <kbd>ctrl</kbd> + <kbd>shift</kbd> + <kbd>+</kbd> : `zoom in`
- <kbd>ctrl</kbd> + <kbd>-</kbd> : `zoom out`

[↑ top](#bash-vim)
<br><br><br><br>
<hr>







#### tmux

- <kbd>ctrl</kbd> + <kbd>b</kbd>, <kbd>t</kbd> : `show time`
- <kbd>ctrl</kbd> + <kbd>b</kbd>, <kbd>%</kbd> : `split vertically`
- <kbd>ctrl</kbd> + <kbd>b</kbd>, <kbd>"</kbd> : `split horizontally`
- <kbd>ctrl</kbd> + <kbd>b</kbd>, <kbd>arrow</kbd> : `switch`
- <kbd>ctrl</kbd> + <kbd>b</kbd> + <kbd>arrow</kbd> : `resize`
- <kbd>ctrl</kbd> + <kbd>b</kbd>, `:resize-pane -R 10` : `resize`
- <kbd>ctrl</kbd> + <kbd>b</kbd>, <kbd>x</kbd> : `kill pane`

[↑ top](#bash-vim)
<br><br><br><br>
<hr>






#### `vim`

```
           l
          /
         k
        /
       j
      /
     h
```

- <kbd>i</kbd>  : `insert mode`
- <kbd>o</kbd>  : `insert a new line and start insert mode`
- <kbd>:</kbd> + <kbd>1</kbd> : `beginning of a file`
- <kbd>H</kbd>  : `beginning of a file`
- <kbd>G</kbd> : `end of a file`
- <kbd>0</kbd> : `beginning of a line`
- <kbd>$</kbd> : `end of a line`
- <kbd>ctrl</kbd> + <kbd>w</kbd>, <kbd>s</kbd> : `split horizontally`
- <kbd>ctrl</kbd> + <kbd>w</kbd>, <kbd>v</kbd> : `split vertically`
- <kbd>ctrl</kbd> + <kbd>w</kbd>, <kbd>arrow</kbd> : `move between panels`
- <kbd>ctrl</kbd> + <kbd>w</kbd>, <kbd>ctrl</kbd> + <kbd>w</kbd> : `move between panels`
- <kbd>ctrl</kbd> + <kbd>w</kbd>, <kbd>q</kbd> : `close a panel`
- <kbd>:vertical resize +10</kbd> : `horizontal increase`
- <kbd>F2</kbd> : `toggle NerdTree`
- <kbd>F3</kbd> : `cycle NerdTree`
- <kbd>F3</kbd>, <kbd>m</kbd>, <kbd>a</kbd> : `create a new file from NerdTree`
- <kbd>shift</kbd> + <kbd>r</kbd> : `refresh NerdTree`
- <kbd>g</kbd> + <kbd>c</kbd> : `toggle comments`
- <kbd>d</kbd> + <kbd>d</kbd> : `delete the current line`
- <kbd>d</kbd> + <kbd>w</kbd> : `delete the next word`
- <kbd>d</kbd> + <kbd>→</kbd> : `delete the character after cursor`
- <kbd>d</kbd> + <kbd>←</kbd> : `delete the character before cursor`
- <kbd>b</kbd> : `beginning of present or previous word`
- <kbd>w</kbd> : `beginning of next word`
- <kbd>e</kbd> : `end of present word`
- <kbd>ea</kbd> : `append at the end of word`
- <kbd>A</kbd> : `append at the end of line`
- <kbd>b</kbd> + <kbd>v</kbd> + <kbd>e</kbd> : `select the current word`
- <kbd>d</kbd> + <kbd>a</kbd> + <kbd>w</kbd> : `delete the current word`
- <kbd>c</kbd> + <kbd>a</kbd> + <kbd>w</kbd> : `delete the current word, and insert`
- <kbd>y</kbd> : `copy the selected characters`
- <kbd>yy</kbd> : `copy the entire line`
- <kbd>“</kbd> + <kbd>+</kbd> + <kbd>y</kbd> : `copy to clipboard`
- <kbd>u</kbd> : `undo`
- <kbd>ctrl</kbd> + <kbd>r</kbd> : `undo the undo`
- <kbd>P</kbd> : `paste before the cursor`
- <kbd>p</kbd> : `paste after`
- <kbd>/</kbd>, <kbd>n</kbd> : `find, and next`
- <kbd>:%s/old/new/gc</kbd> : `find, and replace`
- <kbd>ctrl</kbd> + <kbd>n</kbd>, <kbd>i</kbd> : `expand selection, and edit`
- <kbd>ctrl</kbd> + <kbd>z</kbd> : `go back to terminal`
- <kbd>fg</kbd> : `come back to vim`
- <kbd>ctrl</kbd> + <kbd>p</kbd> : `in insert mode, autocomplete`
- <kbd>ctrl</kbd> + <kbd>p</kbd> : `in visual mode, fuzzy file search`
- <kbd>ctrl</kbd> + <kbd>p</kbd>, <kbd>p</kbd> : `in visual mode, repeat fuzzy file search`
- <kbd>:GoDoc</kbd> : `see the documentation`
- <kbd>:GoPlay</kbd> : `Go playground`
- <kbd>space</kbd> + <kbd>w</kbd> : `:w`
- <kbd>space</kbd> + <kbd>q</kbd> : `:wq`

[↑ top](#bash-vim)
<br><br><br><br>
<hr>
