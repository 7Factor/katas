#!/bin/sh
export GOPATH=$(pwd)
export GOBIN=$GOPATH/bin

cd src && glide install