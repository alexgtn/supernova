/*
Copyright © 2022 Alex Gutan

*/
package main

import (
	"github.com/alexgtn/supernova/cmd"

	_ "github.com/lib/pq"
)

func main() {
	cmd.Execute()
}
