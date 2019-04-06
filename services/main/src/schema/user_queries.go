package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/snapiz/go-vue-starter/packages/cgo"
)

var meQuery = &graphql.Field{
	Type: userType,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		ctx := cgo.FromContext(p.Context)

		if ctx.ID == "" {
			return nil, nil
		}

		return ctx, nil
	},
}
