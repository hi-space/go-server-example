## Dependencies

```sh
go get google.golang.org/grpc
go get go.mongodb.org/mongo-driver 
go get google.golang.org/protobuf
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

```

## References

- [[Tutorial, Part 2] How to develop Go gRPC microservice with HTTP/REST endpoint, middleware, Kubernetes deployment, etc.](https://medium.com/@amsokol.com/tutorial-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-kubernetes-af1fff81aeb2)
  - [Github](https://github.com/amsokol/go-grpc-http-rest-microservice-tutorial/tree/part2)
- [REST over gRPC with grpc-gateway for Go](https://medium.com/swlh/rest-over-grpc-with-grpc-gateway-for-go-9584bfcbb835)