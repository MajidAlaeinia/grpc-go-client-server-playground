# Go GRPC Client Server Playground

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