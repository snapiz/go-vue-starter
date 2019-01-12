package schema

import (
	"github.com/snapiz/go-vue-starter/api/common"
	"github.com/snapiz/go-vue-starter/api/user"

	"github.com/graphql-go/graphql"
)

// Schema defined in init()
var Schema graphql.Schema

func init() {
	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"node": common.NodeDefinitions.NodeField,
				"me":   user.MeQuery,
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"updateUser": user.UpdateUserMutation,
				"changePassword": user.ChangePasswordMutation,
			},
		}),
	})
	if err != nil {
		// panic if there is an error in schema
		panic(err)
	}
}
