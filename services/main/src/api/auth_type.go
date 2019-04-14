package api

import (
	"github.com/graphql-go/graphql"
)

var oAuth2providerType = graphql.NewEnum(graphql.EnumConfig{
	Name:        "OAuth2Provider",
	Description: "OAuth2 provider",
	Values: graphql.EnumValueConfigMap{
		"GOOGLE": &graphql.EnumValueConfig{
			Value: "google",
		},
		"FACEBOOK": &graphql.EnumValueConfig{
			Value: "facebook",
		},
	},
})

var oAuth2Type = graphql.NewObject(graphql.ObjectConfig{
	Name:   "OAuth2",
	Fields: graphql.Fields{
		"google": &graphql.Field{
			Type:        graphql.String,
			Description: "Google OAuth2 redirect URL.",
		},
		"facebook": &graphql.Field{
			Type:        graphql.String,
			Description: "Facebook OAuth2 redirect URL.",
		},
	},
})
