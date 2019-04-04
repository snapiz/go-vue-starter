package main

import (
	"log"

	"github.com/pressly/goose"
	"github.com/snapiz/go-vue-starter/packages/cgo"
	"github.com/spf13/cobra"
)

func up(cmd *cobra.Command, args []string) {
	db, err := cgo.NewDB("", false)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal(err)
	}
}

func init() {
	root.AddCommand(&cobra.Command{
		Use:   "db:up",
		Short: "Migrate database schema",
		Run:   up,
	})
}
