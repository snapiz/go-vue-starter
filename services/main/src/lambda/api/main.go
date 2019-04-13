package main

import (
	"github.com/snapiz/go-vue-starter/services/main/src/api"
	"github.com/snapiz/go-vue-starter/packages/cgo"
)

func main()  {
	cgo.StartLambda(api.AddRoutes)
}