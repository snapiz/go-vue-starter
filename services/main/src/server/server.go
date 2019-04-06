package main

import (
	"github.com/snapiz/go-vue-starter/packages/cgo"
	"github.com/snapiz/go-vue-starter/services/main/src/schema"
	"github.com/snapiz/go-vue-starter/services/main/src/utils"
)

func main() {
	cgo.Start([]cgo.ServiceConfig{
		cgo.ServiceConfig{
			Name:   "main",
			Schema: schema.Schema,
			Before: func(ctx cgo.Context) cgo.Context {
				return utils.NewContextWithUser(ctx)
			},
		},
	})
}
