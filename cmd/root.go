/*
Copyright © 2022 Alex Gutan

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/alexgtn/supernova/config"
)

var cfg *config.Config

var rootCmd = &cobra.Command{
	Use: "supernova",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	cfg = config.FromLocalFile(".", ".env")
}
