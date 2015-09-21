[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# Python: introduction

- [Reference](#reference)
- [Install](#install)
- [Hello World!](#hello-world)

[↑ top](#python-introduction)
<br><br><br><br>
<hr>








#### Reference

- [python.org](https://www.python.org/)
- [Python 2.x documentation](https://docs.python.org/2/)
- [Learn python the hard way](http://learnpythonthehardway.org/book/)

[↑ top](#python-introduction)
<br><br><br><br>
<hr>









#### Install

Please visit [here](https://www.python.org/downloads/).

```bash
#!/bin/bash

cd $HOME;
sudo apt-get -y install python-pip python-dev python-all \
python-psycopg2 python-numpy python-pandas python-mysqldb;

sudo pip install --upgrade pip;
sudo pip install --upgrade psycopg2;
sudo pip install --upgrade pyyaml;
sudo pip install --upgrade gevent;
sudo pip install --upgrade sqlalchemy;
sudo pip install --upgrade boto;

```

[↑ top](#python-introduction)
<br><br><br><br>
<hr>








#### Hello World!

```python
#!/usr/bin/python -u
print "Hello World!"
```

You can either:

- `python code/hello.py`
- `cd code/` and `sudo chmod +x hello.py` and `./hello.py`
- `python` and `print "Hello World!"` in the intepreter

[↑ top](#python-introduction)
<br><br><br><br>
<hr>
