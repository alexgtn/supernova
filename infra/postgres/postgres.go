package postgres

import (
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/alexgtn/supernova/ent"
	"log"
)

// Open new connection
func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}
