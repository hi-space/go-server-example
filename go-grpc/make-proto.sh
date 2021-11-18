#!/bin/sh

protoc -I . proto/helloworld.proto --go_out=plugins=grpc:.
