#!/bin/sh
export GOPATH=$(pwd)
export GOBIN=$GOPATH/bin

go test -v 7factor.io/kata/_unittests