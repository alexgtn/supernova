![](https://s1.gifyu.com/images/nLJguQ9---Imgur.gif)


Run initial setup commands

```
mkdir postgres-data
chmod u+x task atlas-linux-amd64
```

## Start service

`./task main`

## Protobuf

Requirements: protoc-gen-go, protoc-gen-go-grpc

```
./task pb
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

- [ ] Auth/AuthZ with Auth0 and grpc_auth middleware
- [ ] Viper config
- [ ] gRPC request validation https://github.com/envoyproxy/protoc-gen-validate
- [ ] linter
- [ ] Add tests
- [ ] db transactions
- [ ] Gitlab CI