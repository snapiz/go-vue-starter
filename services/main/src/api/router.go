package api

import (
	"github.com/snapiz/go-vue-starter/packages/cgo"
	"github.com/snapiz/go-vue-starter/services/main/src/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// AddRoutes to cgo.Router 
func AddRoutes(r *cgo.Router) {
	r.Add(&cgo.RouteConfig{
		Path:   "/api/main",
		Schema: &Schema,
		FetchUser: func(queryMod qm.QueryMod) (interface{}, error) {
			users, err := models.Users(queryMod).AllG()

			if err != nil {
				return nil, err
			}

			if users == nil {
				return nil, nil
			}

			return users[0], nil
		},
	})
}
