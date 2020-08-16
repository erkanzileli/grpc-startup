# grpc-startup

```
go install \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    github.com/golang/protobuf/protoc-gen-go
```

This will place three binaries in your $GOBIN;

- protoc-gen-grpc-gateway
- protoc-gen-swagger
- protoc-gen-go

Make sure that your $GOBIN is in your $PATH.

To generate proto resources from proto files, run command below

```
$ ./generate
```