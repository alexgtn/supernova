![](https://s1.gifyu.com/images/nLJguQ9---Imgur.gif)


Run initial setup commands

```
mkdir postgres-data
chmod u+x task atlas-linux-amd64
```

## Start service

`./task main`

gRPC gateway

`./task http`

## Protobuf

GRPC Requirements: 
- `protoc-gen-go` `protoc-gen-go-grpc` https://grpc.io/docs/languages/go/quickstart/
- `protoc-gen-validate` https://github.com/envoyproxy/protoc-gen-validate
- grpc gateway https://github.com/grpc-ecosystem/grpc-gateway

```
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest \
    google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
    github.com/envoyproxy/protoc-gen-validate@latest
```

Generate ent db code 
```
./task gen-schema
```

Generate protobuf and API docs
```
./task gen
```

## Migrations
```
./task generate-migration
./task execute-migration
```

Optional 
```
./task validate-migration
./task rehash-migration
```

## TODO

- [ ] Auth/AuthZ with Auth0 and grpc_auth middleware, or Auth https://casdoor.org/ AuthZ https://casbin.org/en/, or https://github.com/ory/keto
- [x] Viper config
- [x] gRPC request validation https://github.com/envoyproxy/protoc-gen-validate
- [x] linter
- [x] db transactions
- [x] Gitlab CI