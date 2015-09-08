#!/bin/bash
# https://github.com/coreos/rkt/blob/master/Documentation/getting-started-guide.md

CGO_ENABLED=0 GOOS=linux go build -o code -a -installsuffix cgo .;
file code;
ldd code;
sudo ./actool --debug validate manifest.json;

mkdir -p code-layout/rootfs;
mkdir -p code-layout/rootfs/bin;

cp manifest.json code-layout/manifest;

cp code code-layout/rootfs/bin;
cp -rf static/ code-layout/rootfs/bin;
cp -rf templates/ code-layout/rootfs/bin;

sudo ./actool build --overwrite code-layout/ code-0.0.1-linux-amd64.aci;
sudo ./actool --debug validate code-0.0.1-linux-amd64.aci;

sudo ./rkt metadata-service  &
sudo ./rkt --insecure-skip-verify run code-0.0.1-linux-amd64.aci;
