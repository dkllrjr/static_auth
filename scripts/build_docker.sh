#!/bin/bash

VERSION="0.1.1"
BASE=`pwd`

cd $BASE/..

docker buildx build --push --platform linux/amd64,linux/arm64 --tag dkllrjr/static_auth:$VERSION .

#docker build --platform linux/amd64 --tag dkllrjr/static_auth:$VERSION-amd64 .
#docker build --platform linux/amd64 --tag dkllrjr/static_auth:latest-amd64 .

#docker build --platform linux/arm64 --tag dkllrjr/static_auth:$VERSION-arm64 .
#docker build --platform linux/arm64 --tag dkllrjr/static_auth:latest-arm64 .

#docker push -a dkllrjr/static_auth
