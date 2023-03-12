#!/bin/bash

BASE=`pwd`

cd $BASE/../src/
source .env
go build
