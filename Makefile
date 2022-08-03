generate:
	buf generate
	docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli generate -i /local/docs/user.swagger.json -l html2 -o /local/docs

http:
	go run main.go http

main:
	go run main.go main

gen-schema:
    # wipe existing generated code
	find tools/ent/codegen -type f \( ! -iname "codegen.go" \) -delete
    # generate code
	go run entgo.io/ent/cmd/ent generate ./tools/ent/schema --target ./tools/ent/codegen

execute-migration:
	go run main.go execute-migration

generate-migration:
	go run main.go generate-migration
	go generate ./tools/ent

validate-migration:
	atlas-linux-amd64 migrate validate migrations

rehash-migration:
	atlas-linux-amd64 migrate hash --force