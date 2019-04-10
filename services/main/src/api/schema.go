package api

import (
	"log"
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
				"node": nodeDefinitions.NodeField,
				"me":   meQuery,
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"updateUser":     updateUserMutation,
				"changePassword": changePasswordMutation,
			},
		}),
	})

	if err != nil {
		log.Fatal(err)
	}
}
