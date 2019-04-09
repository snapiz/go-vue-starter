package main

import (
	"log"

	_ "github.com/snapiz/go-vue-starter/services/main/src/db"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "gvs-main-cli",
	Short: "Go vue start main CLI tool",
}

func main() {
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
