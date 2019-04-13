package main

import (
	"github.com/snapiz/go-vue-starter/packages/cgo"
	"github.com/snapiz/go-vue-starter/services/main/src/auth"
)

func main() {
	cgo.StartLambda(auth.AddRoutes)
}
