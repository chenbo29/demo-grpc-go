# demo-grpc-go

## demo base from grpc.io

[grpc golang quickstart](https://grpc.io/docs/languages/go/quickstart/)

## client

```shell
    docker exec -it demo-grpc-go_rpc_1 go run greeter_client/main.go --name=chenbo
```

## compile .proto files

```shell
protoc --go_out=. --go_opt=paths=source_relative \ 
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
helloworld/helloworld.proto

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protos/helloworld/helloworld.proto

```
