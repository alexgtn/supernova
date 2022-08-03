gen-buf:
	buf generate
	docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli generate -i /local/docs/user.swagger.json -l html2 -o /local/docs

http:
	go run main.go http

main:
	go run main.go main

gen-db-schema:
    # wipe existing generated code
	find internal/infra/ent/gen -type f \( ! -iname "codegen.go" \) -delete
    # generate code
	go run entgo.io/ent/cmd/ent generate ./internal/infra/ent/schema --target ./internal/infra/ent/gen

execute-migration:
	go run main.go execute-migration

generate-migration:
	go run main.go generate-migration
	go generate ./tools/ent

validate-migration:
	atlas-linux-amd64 migrate validate migrations

rehash-migration:
	atlas-linux-amd64 migrate hash --force