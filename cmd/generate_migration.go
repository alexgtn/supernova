/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/sqltool"
	entsql "entgo.io/ent/dialect/sql"

	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"

	"github.com/spf13/cobra"

	"github.com/alexgtn/supernova/internal/infra/postgres"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "generate-migration",
	Short: "Generate new migration from schema updates",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generating new migration")
		// Load the graph.
		graph, err := entc.LoadGraph("./ent/schema", &gen.Config{})
		if err != nil {
			log.Fatalln(err)
		}
		tbls, err := graph.Tables()
		if err != nil {
			log.Fatalln(err)
		}
		// Create a local migration directory.
		d, err := migrate.NewLocalDir("tools/migrations")
		if err != nil {
			log.Fatalln(err)
		}
		client := postgres.Open(cfg.DatabaseURL)
		defer func(client *entsql.Driver) {
			err := client.Close()
			if err != nil {
				log.Fatal("error closing client")
			}
		}(client)

		// Inspect it and compare it with the graph.
		m, err := schema.NewMigrate(client, schema.WithDir(d),
			schema.WithFormatter(sqltool.GolangMigrateFormatter),
			schema.WithSumFile(),
		)
		if err != nil {
			log.Fatalln(err)
		}
		if err := m.Diff(context.Background(), tbls...); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
