package postgres

import (
	"database/sql"
	"github.com/alexgtn/supernova/ent"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
)

// OpenEnt new connection
func OpenEnt(databaseUrl string) *ent.Client {
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

// Open new connection
func Open(databaseUrl string) *entsql.Driver {
	db, err := entsql.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
