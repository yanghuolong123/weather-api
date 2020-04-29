#!/bin/bash

PROJECT="app"
DIR=`pwd`
GOPATH="$DIR/code"

export GOPATH="$GOPATH"
echo "gopath:"$GOPATH

cd code

export GOPROXY=https://goproxy.io

go get -u github.com/astaxie/beego
#go get -u github.com/beego/bee


export PATH="$GOPATH/bin:$PATH"
echo "path:"$PATH


#bee new $PROJECT
bee api $PROJECT
cd src/$PROJECT
go mod init $PROJECT


echo "init finished..."
