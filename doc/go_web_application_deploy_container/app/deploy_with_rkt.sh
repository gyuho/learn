#!/bin/bash
# https://github.com/coreos/rkt/blob/master/Documentation/getting-started-guide.md

CGO_ENABLED=0 GOOS=linux go build -o code -a -installsuffix cgo .;
file code;
ldd code;
sudo ./actool --debug validate manifest.json;

mkdir -p image/rootfs/usr/bin;

sudo cp manifest.json image/manifest;

sudo cp code image/rootfs/usr/bin;
sudo cp -rf static/ image/rootfs/usr/bin;
sudo cp -rf templates/ image/rootfs/usr/bin;

sudo ./actool build --overwrite image/ code-0.0.1-linux-amd64.aci;
sudo ./actool --debug validate code-0.0.1-linux-amd64.aci;

sudo ./rkt metadata-service >/dev/null 2>&1 & # run in background

sudo ./rkt --insecure-skip-verify run \
code-0.0.1-linux-amd64.aci \
--volume static,kind=host,source=/usr/bin/static \
--volume templates,kind=host,source=/usr/bin/templates \
-- \
;
