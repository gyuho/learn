[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# Go: introduction

- [Reference](#reference)
- [Install](#install)
- [Hello World!](#hello-world)

[↑ top](#go-introduction)
<br><br><br><br>
<hr>







#### Reference

- [golang.org](http://golang.org/)
- [Go Tour](http://tour.golang.org/welcome/1)

[↑ top](#go-introduction)
<br><br><br><br>
<hr>









#### Install

Please visit [here](http://golang.org/doc/install).
In `Ubuntu 14.04.3 LTS` (Linux Debian distribution), you can run:

```bash
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
```

[↑ top](#go-introduction)
<br><br><br><br>
<hr>









#### Hello World!

Try this [code](http://play.golang.org/p/OccSs5jC9Y):

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

You can either:

- `go run hello/main.go`
- `cd hello/` and `go build` and `./hello`
- `cd hello/` and `go install` and `hello`

[↑ top](#go-introduction)
<br><br><br><br>
<hr>
