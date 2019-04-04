package main

import (
	"log"

	"github.com/pressly/goose"
	"github.com/spf13/cobra"
)

func init()  {
	root.AddCommand(&cobra.Command{
		Use:   "db:migration",
		Short: "Generate SQL migration file",
		Run: func(cmd *cobra.Command, args []string) {
			err := goose.Create(nil, "migrations", args[0], "sql")
			if err != nil {
				log.Fatal(err)
			}
		},
	})
}
