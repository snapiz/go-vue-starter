package main

import (
	"github.com/snapiz/go-vue-starter/packages/cgo"
	"github.com/snapiz/go-vue-starter/services/main/src/api"
	"github.com/snapiz/go-vue-starter/services/main/src/auth"
	_ "github.com/snapiz/go-vue-starter/services/main/src/db"
)

func main() {
	cgo.Start(func(r *cgo.Router) {
		api.AddRoutes(r)
		auth.AddRoutes(r)
	})
}
