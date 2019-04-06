package cgo

import (
	"go/build"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
}

func loadEnv() {
	env := os.Getenv("GO_ENV")

	if env == "" {
		return
	}

	godotenv.Load(".env." + env)
	err := godotenv.Load()

	if env == "test" && err != nil {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		parent := filepath.Dir(wd)
		os.Chdir(parent)
		loadEnv()
	}
}

// GetGOPath Get $GOPATH or default build
func GetGOPath() string {
	gopath := os.Getenv("GOPATH")

	if gopath == "" {
		return build.Default.GOPATH
	}

	return gopath
}
