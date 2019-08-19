#!/usr/bin/env bash

#echo 'Linux' && gox -osarch="linux/amd64" --output="build/app"
echo 'Windows' && gox -osarch="windows/amd64" --output="build/app"

#chmod 755 Dockerfile

docker build -t azurewtgo .

## See instruction: https://cloud.docker.com/repository/docker/
# docker tag local-image:tagname reponame:tagname
docker tag azurewtgo:latest ilovejs/idk:latest

#docker push reponame:tagname
docker push ilovejs/idk:latest

#netstat -ant | grep 8585

#docker run -it --publish 1459:8585 ilovejs/idk

