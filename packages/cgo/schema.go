package cgo

import (
	"fmt"
	"strings"

	"github.com/graphql-go/graphql"
)

// UserRoleType role of user
var UserRoleType = graphql.NewEnum(graphql.EnumConfig{
	Name:        "UserRole",
	Description: "The role of the user",
	Values: graphql.EnumValueConfigMap{
		strings.ToUpper(ContextRoleAdmin): &graphql.EnumValueConfig{
			Value: ContextRoleAdmin,
		},
		strings.ToUpper(ContextRoleStaff): &graphql.EnumValueConfig{
			Value: ContextRoleStaff,
		},
		strings.ToUpper(ContextRoleUser): &graphql.EnumValueConfig{
			Value: ContextRoleUser,
		},
	},
})

// UserStateType state of user
var UserStateType = graphql.NewEnum(graphql.EnumConfig{
	Name:        "UserState",
	Description: "The state of the user",
	Values: graphql.EnumValueConfigMap{
		strings.ToUpper(ContextStateEnable): &graphql.EnumValueConfig{
			Value: ContextStateEnable,
		},
		strings.ToUpper(ContextStateDisable): &graphql.EnumValueConfig{
			Value: ContextStateDisable,
		},
		strings.ToUpper(ContextStateMaintenance): &graphql.EnumValueConfig{
			Value: ContextStateMaintenance,
		},
	},
})

// NewUserTypeFields return default UserType fields
func NewUserTypeFields(fields graphql.Fields) graphql.Fields {
	userFields := graphql.Fields{
		"id": GlobalIDField("User"),
		"email": &graphql.Field{
			Type:        graphql.String,
			Description: "The email of the user.",
		},
		"username": &graphql.Field{
			Type:        graphql.String,
			Description: "The name of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u := p.Source.(Context)
				return u.Username.Ptr(), nil
			},
		},
		"hasPassword": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Is password is defined.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u := p.Source.(Context)
				return u.Password.Ptr() != nil, nil
			},
		},
		"displayName": &graphql.Field{
			Type:        graphql.String,
			Description: "The name of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u := p.Source.(Context)
				return u.DisplayName.Ptr(), nil
			},
		},
		"picture": &graphql.Field{
			Type:        graphql.String,
			Description: "The picture of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u := p.Source.(Context)
				return u.Picture.Ptr(), nil
			},
		},
		"avatar": &graphql.Field{
			Type:        graphql.String,
			Description: "The avatar of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				u := p.Source.(Context)
				if u.Picture.Ptr() == nil {
					return fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon&s=68", u.EmailHash), nil
				}

				return u.Picture.Ptr(), nil
			},
		},
		"role": &graphql.Field{
			Type: UserRoleType,
		},
		"state": &graphql.Field{
			Type: UserStateType,
		},
	}

	for key, field := range fields {
		userFields[key] = field
	}

	return userFields
}
