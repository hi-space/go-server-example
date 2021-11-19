## Docker image build

```sh
docker image build -t go_server:latest .
```

## Docker run

```sh
docker container run -t -p 9000:8080 go_server:latest
```

Docker 내부의 8080 포트와 Host OS의 9000 포트를 연결해서 실행

## Test

```sh
curl http://localhost:9000
```

localhost 9000에 curl 을 해보면 서버가 정상적으로 떠있는 것을 볼 수 있다.

# References

- [https://dydtjr1128.gitbook.io/understanding-docker/2.release-docker-container/1-make-simple-docker-image](https://dydtjr1128.gitbook.io/understanding-docker/2.release-docker-container/1-make-simple-docker-image)
