package api

import "github.com/graphql-go/graphql"

var meQuery = &graphql.Field{
	Type: userType,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		c := NewContext(p.Context)

		return c.User, nil
	},
}
