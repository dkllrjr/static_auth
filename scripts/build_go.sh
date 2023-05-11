#!/bin/bash

BASE=`pwd`

cd $BASE/../src/
ln -s $BASE/.env .
source .env

go build
rm .env
