#!/bin/bash
# https://github.com/coreos/rkt/blob/master/Documentation/getting-started-guide.md

printf "Compiling Go\n"
CGO_ENABLED=0 GOOS=linux go build -o app -a -installsuffix cgo .;
file app;
ldd app;
sudo ./actool --debug validate manifest.json;

printf "Copying files\n"
mkdir -p image/rootfs/usr/bin;
sudo cp manifest.json image/manifest;
sudo cp app image/rootfs/usr/bin;
sudo cp -rf static/ image/rootfs/usr/bin;
sudo cp -rf templates/ image/rootfs/usr/bin;

printf "Building ACI with acitool\n"
sudo ./actool build --overwrite image/ app.aci;
sudo ./actool --debug validate app.aci;

printf "Running rkt in background\n"
# sudo ./rkt metadata-service >/dev/null 2>&1 &

printf "Executing with rkt\n"
sudo ./rkt --insecure-skip-verify run \
app.aci \
--volume static,kind=host,source=/usr/bin/static \
--volume templates,kind=host,source=/usr/bin/templates \
-- \
;
