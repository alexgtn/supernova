package main

import (
	"context"
	"github.com/alexgtn/supernova/infra/postgres"
	"log"

	"entgo.io/ent/dialect/sql/schema"
)

func main() {
	client := postgres.Open("postgresql://default:default@localhost:5432/postgres")
	defer client.Close()
	ctx := context.Background()
	// Run migration.
	err := client.Schema.Create(ctx, schema.WithAtlas(true))
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
