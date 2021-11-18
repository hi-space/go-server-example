all: protoc client server

protoc:
	@echo "Generating Go files"
	cd proto && protoc --go_out=plugins=grpc:. *.proto

server: protoc
	@echo "Building server"
	go build -o server \
		github.com/hi-space/go-server-example/tree/main/go-grpc-stream/src/server

client: protoc
	@echo "Building client"
	go build -o client \
		github.com/hi-space/go-server-example/tree/main/go-grpc-stream/src/client

.PHONY: client server protoc