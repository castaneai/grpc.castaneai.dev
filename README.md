# grpc.castaneai.dev

- A Minimal gRPC echo server with bidirectional streaming RPC
- [Server Reflection](https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md) support
- service proto: [proto/echo.proto](./proto/echo.proto)
- YAML definition for Cloud Run

## Usage

```sh
# via grpcurl
grpcurl -d '{"message": "hello"}' grpc.castaneai.dev:443 EchoService/StreamingEcho

# A testclient creates multiple streams with EchoService/StreamingEcho
go run testclient/testclient.go grpc.castaneai.dev:443 3
```

## Deploying to Cloud Run

### requirements

- Google Cloud SDK
- [ko](https://github.com/google/ko)

```sh
export KO_DOCKER_REPO=gcr.io/<YOUR_GCP_PROJECT>
make deploy
```

## LICENSE

MIT
