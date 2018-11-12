#!/bin/sh
export GOPATH=$(pwd)
export GOBIN=$GOPATH/bin

ginkgo watch -r src/7factor.io/kata/_unittests