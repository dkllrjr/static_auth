#!/bin/bash

BASE=`pwd`

cd $BASE/..

docker build -t dkllrjr/static_auth:0.1.1 .
