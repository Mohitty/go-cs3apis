#!/bin/bash

GOPATH=`go env GOPATH`
PATH=$PATH:$GOPATH/bin
rm -rf cs3/*
rm -rf build
git clone https://github.com/cs3org/cs3apis build && cd build && make deps && make && cd ..
cp  prototool_gen_go.yaml build/prototool.yaml
#go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
#go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
#go get -u github.com/golang/protobuf/protoc-gen-go
cd build && prototool generate && cd ..
rm -rf build

# change import paths
egrep -lR 'github.com/cs3org/cs3/' ./cs3 | xargs sed -i 's|github.com/cs3org/cs3/|github.com/cs3org/go-cs3apis/cs3/|g'
