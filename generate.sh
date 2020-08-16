#! /bin/bash

set -e

PROTO_FILES=$(find . -type f -name "*.proto" | xargs)

for proto_file in $PROTO_FILES ; do
  protoc -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --proto_path=$PWD --go_out=plugins=grpc:. $proto_file
  protoc -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --proto_path=$PWD --grpc-gateway_out=logtostderr=true,paths=source_relative:. $proto_file
  protoc -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --proto_path=$PWD --swagger_out=grpc_api_configuration:. $proto_file
  echo "Generated code for $proto_file"
done
