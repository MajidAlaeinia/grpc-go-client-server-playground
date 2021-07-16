# Go gRPC Client Server Playground

1. 
```shell
go mod download
```

2.
```shell
go mod vendor
```

3.
```shell
protoc --go_out=plugins=grpc:chat chat.proto
```
4.
```shell
go run server.go
```

5.
```shell
go run client.go
```
