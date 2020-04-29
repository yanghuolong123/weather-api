#!/bin/bash

PROJECT="app"
DIR=`pwd`
GOPATH="$DIR/code"

export GOPATH="$GOPATH"
echo "gopath:"$GOPATH

export PATH="$GOPATH/bin:$PATH"
echo "path:"$PATH

cd code/src/$PROJECT
#bee run 
bee run -gendoc=true -downdoc=true

echo "runing ..."
