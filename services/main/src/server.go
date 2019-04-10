package main

import (
	"github.com/snapiz/go-vue-starter/packages/cgo"
	"github.com/snapiz/go-vue-starter/services/main/src/api"
	"github.com/snapiz/go-vue-starter/services/main/src/auth"
	_ "github.com/snapiz/go-vue-starter/services/main/src/db"
	"github.com/snapiz/go-vue-starter/services/main/src/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func main() {
	cgo.Start("/main", func(queryMod qm.QueryMod) (interface{}, error) {
		users, err := models.Users(queryMod).AllG()

		if err != nil {
			return nil, err
		}

		if users == nil {
			return nil, nil
		}

		return users[0], nil
	}, func(r *cgo.Router) {
		r.Add(&cgo.RouteConfig{
			Path:   "/api",
			Schema: &api.Schema,
		})

		r.Add(&cgo.RouteConfig{
			Path:        "/auth/local",
			HandlerFunc: auth.LocalHandler,
		})

		r.Add(&cgo.RouteConfig{
			Path:        "/auth/local/new",
			HandlerFunc: auth.LocalNewHandler,
		})

		r.Add(&cgo.RouteConfig{
			Path:        "/auth/logout",
			HandlerFunc: auth.LogoutHandler,
		})
	})
}
