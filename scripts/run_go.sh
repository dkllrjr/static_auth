#!/bin/bash

BASE=`pwd`

cd $BASE/../src/
ln -s $BASE/.env .
ln -s $BASE/../website

source .env
go run .

rm .env
rm website
