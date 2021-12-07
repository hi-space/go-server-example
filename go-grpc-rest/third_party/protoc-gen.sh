#!/bin/sh

protoc --proto_path=api/proto --proto_path=third_party --go_out=plugins=grpc:. todo.proto
protoc --proto_path=api/proto --proto_path=third_party --grpc-gateway_out=logtostderr=true:. todo.proto
protoc --proto_path=api/proto --proto_path=third_party --swagger_out=logtostderr=true:api/swagger todo.proto

# protoc --go_out=plugins=grpc:. api/proto/*.proto

# protoc \
# -I. \
# -I third_party \
# --go_out=plugins=grpc:. \
# --grpc-gateway_out=logtostderr=true:. \
# --swagger_out=logtostderr=true:api/swagger/v1 \
# api/proto/v1/todo.proto