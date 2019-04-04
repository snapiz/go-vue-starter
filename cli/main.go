package main

import (
	"log"

	_ "github.com/lib/pq"
	_ "github.com/snapiz/go-vue-starter/packages/cgo"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "gvs-cli",
	Short: "Go vue start CLI tool",
}

func main() {
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
