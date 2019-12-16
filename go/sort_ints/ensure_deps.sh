#!/usr/bin/env bash

export GOPATH=$PWD

cd src/7factor.io && dep ensure
