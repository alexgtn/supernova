/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/sqltool"
	"context"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate called")
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
		d, err := migrate.NewLocalDir("migrations")
		if err != nil {
			log.Fatalln(err)
		}
		// OpenEnt connection to the database.
		//client := postgres.Open("postgresql://default:default@localhost:5432/postgres")
		client, err := sql.Open("postgres", "postgresql://default:default@localhost:5432/postgres?sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}

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
		// You can use the following method to give the migration files a name.
		// if err := m.NamedDiff(context.Background(), "migration_name", tbls...); err != nil {
		//     log.Fatalln(err)
		// }
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
