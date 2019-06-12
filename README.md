# Golang With gRPC

## Dep:

```bash
    go get -u -v github.com/golang/protobuf

    go get -u -v github.com/golang/protobuf/proto

    go get -u -v github.com/gin-gonic/gin
        
    go get -u -v google.golang.org/grpc
```

## Env:

> export PATH=$PATH:$GOPATH/bin

## Build:

**compile proto buffers**

```bash

    cd grpc-crud/server/pb

    protoc *.proto --go_out=plugins=grpc:messages

```

**run server**

```bash
    go run server/exec/main.go
```

**run client**

```bash
    go run client.go
```

### Host

> `http://localhost:8080`

### Endpoints

```
POST   /add/user                 
GET    /find/user/:uid           
PUT    /update/user/:uid         
DELETE /delete/user/:uid         
```