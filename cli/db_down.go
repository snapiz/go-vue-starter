package main

import (
	"log"

	"github.com/pressly/goose"
	"github.com/snapiz/go-vue-starter/packages/cgo"
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(&cobra.Command{
		Use:   "db:down",
		Short: "Rollback database schema",
		Run: func(cmd *cobra.Command, args []string) {
			db, err := cgo.NewDB("", false)

			if err != nil {
				log.Fatal(err)
			}

			defer db.Close()

			if err := goose.Down(db, "migrations"); err != nil {
				log.Fatal(err)
			}
		},
	})
}
