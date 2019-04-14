package api

import (
	"github.com/graphql-go/graphql"
	"github.com/snapiz/go-vue-starter/packages/cgo"
	"github.com/snapiz/go-vue-starter/services/main/src/utils"
)

var oAuth2Query = &graphql.Field{
	Type: oAuth2Type,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		ctx := cgo.FromContext(p.Context)

		return utils.GetProviderAuthURLs(&ctx), nil
	},
}
