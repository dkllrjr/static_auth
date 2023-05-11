#!/bin/bash

BASE=`pwd`

cd $BASE/..

ARCHS=('linux/arm64' 'linux/amd64')

for arch in "${ARCHS[@]}"; do
    docker buildx build --platform $arch --tag dkllrjr/static_auth:0.1.1 .
done
