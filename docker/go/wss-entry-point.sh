#!/bin/bash

app_dir="/app"
cd $app_dir

build_name="$app_dir/bin/srv/server"
cert_file="$app_dir/docker/tls/cert.pem"
key_file="$app_dir/docker/tls/key.pem"
mod_file="$app_dir/go.mod"

if [ "$DEVELOPMENT" = "true" ]; then
  echo "'DEVELOPMENT' mode enabled"
  if [ ! -f "$cert_file" ]  && [ ! -f "$key_file" ]; then
    echo "Cert file not present. Creating ..."
    cd $app_dir/docker/tls

    go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=$VIRTUAL_HOST
    chown $HOST_UID:$HOST_GID cert.pem key.pem

    echo "Cert files created"
    cd $app_dir
  fi
fi

if [ ! -f $mod_file ]; then
  echo "'go.mod' file not present. Initializing module '$MODULE_PATH'"
  go mod init $MODULE_PATH
  chown $HOST_UID:$HOST_GID go.mod
  echo "Module '$MODULE_PATH' initialized."
  echo "Downloading module dependencies ..."
  go mod tidy
  chown $HOST_UID:$HOST_GID go.sum
  echo "Module dependencies downloaded."
fi

CompileDaemon --build="$app_dir/make" --directory="$app_dir" --command=$build_name --exclude-dir=".git" \
  --exclude-dir=".idea"
