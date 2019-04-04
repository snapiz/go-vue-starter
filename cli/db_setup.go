package main

import (
	"fmt"
	"log"

	"github.com/snapiz/go-vue-starter/packages/cgo"
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(&cobra.Command{
		Use:   "db:setup",
		Short: "Create database schema and migrate",
		Run: func(cmd *cobra.Command, args []string) {
			db, err := cgo.NewDB("", true)
			dbname := cgo.GetDBName()

			if err != nil {
				log.Fatal(err)
			}

			defer db.Close()

			db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbname))
			up(root, args)
		},
	})
}
