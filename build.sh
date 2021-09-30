#!/usr/bin/env bash
mkdir -p bin
pushd psynder
go build -o ../bin/server psynder/cmd/server
popd
