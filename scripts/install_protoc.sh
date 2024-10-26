#!/usr/bin/env bash
SOURCE="https://github.com/protocolbuffers/protobuf/releases/download/v23.2/protoc-23.2-linux-x86_64.zip"
FILENAME="third_party/protoc.zip"
wget -O $FILENAME $SOURCE
unzip -o -d third_party/protoc $FILENAME
rm -f $FILENAME