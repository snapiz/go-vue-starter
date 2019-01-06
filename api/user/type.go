package user

import (
	"strings"
	"time"

	"github.com/snapiz/go-vue-starter/api/common"
	"github.com/snapiz/go-vue-starter/db/models"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

// UserRoleType role of user
var UserRoleType = graphql.NewEnum(graphql.EnumConfig{
	Name:        "UserRole",
	Description: "The role of the user",
	Values: graphql.EnumValueConfigMap{
		strings.ToUpper(models.UserRoleAdmin): &graphql.EnumValueConfig{
			Value: models.UserRoleAdmin,
		},
		strings.ToUpper(models.UserRoleStaff): &graphql.EnumValueConfig{
			Value: models.UserRoleStaff,
		},
		strings.ToUpper(models.UserRoleUser): &graphql.EnumValueConfig{
			Value: models.UserRoleUser,
		},
		strings.ToUpper(models.UserRoleDisable): &graphql.EnumValueConfig{
			Value: models.UserRoleDisable,
		},
	},
})

// UserType graphql user type
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": relay.GlobalIDField("User", nil),
		"email": &graphql.Field{
			Type:        graphql.String,
			Description: "The email of the user.",
		},
		"displayName": &graphql.Field{
			Type:        graphql.String,
			Description: "The name of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u := p.Source.(*models.User)
				return u.DisplayName.Ptr(), nil
			},
		},
		"picture": &graphql.Field{
			Type:        graphql.String,
			Description: "The picture of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u := p.Source.(*models.User)
				return u.Picture.Ptr(), nil
			},
		},
		"role": &graphql.Field{
			Type: UserRoleType,
		},
		"createdAt": &graphql.Field{
			Type:        graphql.String,
			Description: "The time of the user was created.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u := p.Source.(*models.User)
				return u.CreatedAt.Format(time.RFC3339), nil
			},
		},
		"updatedAt": &graphql.Field{
			Type:        graphql.String,
			Description: "The time of the user was updated.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u := p.Source.(*models.User)

				if u.UpdatedAt.Ptr() == nil {
					return nil, nil
				}

				updatedAt := *u.UpdatedAt.Ptr()

				return updatedAt.Format(time.RFC3339), nil
			},
		},
	},
	Interfaces: []*graphql.Interface{
		common.NodeDefinitions.NodeInterface,
	},
})
