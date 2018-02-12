#!/bin/sh

set -e -x

PACKAGE_NAME=github.com/Infra-Red/myawesomeapi
PACKAGE_DIR=$GOPATH/src/$PACKAGE_NAME
mkdir -p $PACKAGE_DIR
cp -R git/* $PACKAGE_DIR

go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega
go get github.com/gorilla/mux

ginkgo $GOPATH/src/$PACKAGE_NAME/movie
