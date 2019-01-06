package user

import (
	"github.com/snapiz/go-vue-starter/api/common"

	"github.com/graphql-go/graphql"
)

// MeQuery query of current user
var MeQuery = &graphql.Field{
	Type: UserType,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		c := common.NewContext(p.Context)

		return c.User, nil
	},
}
