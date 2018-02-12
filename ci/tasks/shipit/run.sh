#!/bin/bash

set -e -x

OUTPUT_DIR=$(pwd)/github-release
PACKAGE_NAME=github.com/Infra-Red/myawesomeapi
PACKAGE_DIR=$GOPATH/src/$PACKAGE_NAME
mkdir -p $PACKAGE_DIR
cp -R git/* $PACKAGE_DIR

go get github.com/gorilla/mux

git config --global user.name "Andrei Krasnitski"
git config --global user.email "andrei.krasnitski@altoros.com"

mkdir -p $OUTPUT_DIR/artifacts
env GOOS=linux GOARCH=amd64 go build -o "$OUTPUT_DIR/artifacts/release-binary-linux" $PACKAGE_NAME
env GOOS=darwin GOARCH=amd64 go build -o "$OUTPUT_DIR/artifacts/release-binary-darwin" $PACKAGE_NAME
env GOOS=windows GOARCH=amd64 go build -o "$OUTPUT_DIR/artifacts/release-binary-windows.exe" $PACKAGE_NAME

VERSION=$(date '+%Y%m%d%H%M%S')

echo "$VERSION" > $OUTPUT_DIR/name
echo "$VERSION" > $OUTPUT_DIR/tag
echo "Version $VERSION" > $OUTPUT_DIR/body
