#!/usr/bin/env bash

gox -osarch="linux/amd64" --output="build/app"

chmod 755 Dockerfile
docker build -t durgaprasad-budhwani/azurego .

netstat -ant | grep 8585

docker run --publish 1459:8585 durgaprasad-budhwani/azurego
