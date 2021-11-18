# go-grpc-stream

Simple grpc server streaming capability with Golang, this code will have server and client each:

- Server: Will have a function that will stream 5 responses with slight delay each.
- Client: Will send a request to server and wait on all the responses.

## Started

Run with make file

```sh
make all
```

## Usage

### Build Docker Image using docker compose

```sh
docker-compose -f docker-compose.yml build
```

### Run up all docker containers using docker compose

```sh
docker-compose -f docker-compose.yml up
```

# References

- [https://www.freecodecamp.org/news/grpc-server-side-streaming-with-go/](https://www.freecodecamp.org/news/grpc-server-side-streaming-with-go/)
- [https://github.com/pramonow/go-grpc-server-streaming-example](https://github.com/pramonow/go-grpc-server-streaming-example)
