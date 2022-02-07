# grpc.castaneai.dev

## Usage

```sh
# via grpcurl
grpcurl -d '{"message": "hello"}' grpc.castaneai.dev:443 EchoService/StreamingEcho

# via test-client
go run testclient/testclient.go grpc.castaneai.dev:443
```

## Deploy

### requirements

- [ko](https://github.com/google/ko)

```sh
export KO_DOCKER_REPO=asia.gcr.io/<GCP_PROJECT>
make deploy
```
