#!/bin/bash

app_dir="/app"
build_dir="$app_dir/bin/srv"
build_name="$build_dir/server"

rm -rf $build_dir
mkdir -p $build_dir

# -buildvcs=false  - SEE: https://github.com/golang/go/issues/51253
go build --buildvcs=false -o $build_dir/server $app_dir/application
chown $HOST_UID:$HOST_GID $build_name

cp $app_dir/config/app.yml $build_dir/config.yml
cp $app_dir/.env $build_dir/.env

cp $app_dir/docker/tls/cert.pem $build_dir/cert.pem
cp $app_dir/docker/tls/key.pem $build_dir/key.pem
cp -R $app_dir/ui $build_dir

chown -R $HOST_UID:$HOST_GID  $build_dir

