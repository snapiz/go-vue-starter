package cgo

import (
	"go/build"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	env := os.Getenv("GO_ENV")

	if env == "" {
		return
	}

	godotenv.Load(".env." + env)
	godotenv.Load()
}

// GetGOPath Get $GOPATH or default build
func GetGOPath() string {
	gopath := os.Getenv("GOPATH")

	if gopath == "" {
		return build.Default.GOPATH
	}

	return gopath
}
