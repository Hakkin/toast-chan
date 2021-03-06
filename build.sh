#!/bin/bash

GOOS="linux"
GOARCH="amd64"
ldflags="-s -w"
output="./build/"
binary="toastchan"

export GOOS="$GOOS"
export GOARCH="$GOARCH"

go build -v -o ""$output"/$binary" -ldflags "$ldflags"
cp *.json "$output"