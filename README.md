<p align="center">
  <img width="400" src="gopher.png" />
  <br />
  Kudos to https://gopherize.me/ for the cute logo
</p>

## Supernova

This repo provides:
- boilerplate code based on best practices for kickstarting go projects
- clean architecture based on [hexagonal architecture](https://herbertograca.com/2017/11/16/explicit-architecture-01-ddd-hexagonal-onion-clean-cqrs-how-i-put-it-all-together/)
- gRPC API, Protobuf, HTTP gateway, codegen with [buf](https://github.com/bufbuild/buf), API documentation, [go-grpc-middleware](https://github.com/grpc-ecosystem/go-grpc-middleware) logging w/ [zap](https://github.com/uber-go/zap), message validation 
- [Cobra](https://github.com/spf13/cobra) and [Viper](https://github.com/spf13/viper) integration
- [golang-ci](https://github.com/golangci/golangci-lint) linter
- CI config for Gitlab (wake me up when september.. when Github Actions catches up)
- Docker & docker-compose: [traefik](https://github.com/traefik/traefik) proxy, Postgres w/ auto-backup to S3, Datadog metrics, [Watchtower](https://github.com/containrrr/watchtower) auto-deploy 
- [ent](https://github.com/ent/ent) ORM, codegen, migrations


## Start service

```
make main
```

gRPC gateway

```
make http
```

## Protobuf

GRPC Requirements: 
- `protoc-gen-go` `protoc-gen-go-grpc` https://grpc.io/docs/languages/go/quickstart/
- `protoc-gen-validate` https://github.com/envoyproxy/protoc-gen-validate
- grpc gateway https://github.com/grpc-ecosystem/grpc-gateway

```
go install \
    github.com/bufbuild/buf/cmd/buf@v1.7.0 \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest \
    google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
    github.com/envoyproxy/protoc-gen-validate@latest
```

Generate ent db code 
```
make gen-schema
```

Generate protobuf and API docs
```
make gen-buf
```

## Migrations
```
make generate-migration
make execute-migration
```

Optional 
```
make validate-migration
make rehash-migration
```

## TODO

- [ ] Cobra CMD for prod/dev env
- [ ] Auth/AuthZ with Auth0 and grpc_auth middleware, or Auth https://casdoor.org/ AuthZ https://casbin.org/en/, or https://github.com/ory/keto
- [x] Viper config
- [x] gRPC request validation https://github.com/envoyproxy/protoc-gen-validate
- [x] linter
- [x] db transactions
- [x] Gitlab CI

<p align="center">
  <img src="https://s1.gifyu.com/images/nLJguQ9---Imgur.gif" />
  <br />
  Oasis - Champagne Supernova
</p>