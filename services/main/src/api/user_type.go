package api

import (
	"github.com/graphql-go/graphql"
	"github.com/snapiz/go-vue-starter/packages/cgo"
)

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name:   "User",
	Fields: cgo.NewUserTypeFields(graphql.Fields{}),
	Interfaces: []*graphql.Interface{
		nodeDefinitions.NodeInterface,
	},
})
