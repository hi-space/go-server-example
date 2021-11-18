#!/bin/sh

# protoc -I . proto/helloworld.proto --go_out=plugins=grpc:.
protoc --go_out=plugins=grpc:. src/proto/*.proto
