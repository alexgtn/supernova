/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect/sql/schema"

	"github.com/alexgtn/supernova/ent"
	"github.com/alexgtn/supernova/ent/migrate"
	"github.com/alexgtn/supernova/infra/postgres"

	"github.com/spf13/cobra"
)

// executeMigrationCmd represents the executeMigration command
var executeMigrationCmd = &cobra.Command{
	Use:   "execute-migration",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing migration")
		client := postgres.OpenEnt(cfg.DatabaseURL)

		defer func(client *ent.Client) {
			err := client.Close()
			if err != nil {
				log.Fatal("error closing client")
			}
		}(client)
		ctx := context.Background()
		// Run migration.
		err := client.Schema.Create(ctx,
			schema.WithAtlas(true),
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true))
		if err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(executeMigrationCmd)
}
