Main parts of this code is copied from somewhere on the Internet that I don't recall where it was now. This is just a place to know how to scaffold the client server to be ready to talk to each other.  
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
