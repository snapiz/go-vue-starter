package api

import (
	"context"
	"net/http"
	"time"

	"github.com/snapiz/go-vue-starter/server/auth"
	"github.com/snapiz/go-vue-starter/server/db/models"
	"github.com/snapiz/go-vue-starter/server/db/services"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	validator "gopkg.in/go-playground/validator.v9"
)

var updateUserMutation = relay.MutationWithClientMutationID(relay.MutationConfig{
	Name: "UpdateUser",
	InputFields: graphql.InputObjectConfigFieldMap{
		"displayName": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"picture": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	OutputFields: graphql.Fields{
		"user": &graphql.Field{
			Type: userType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if payload, ok := p.Source.(map[string]interface{}); ok {
					return payload["user"].(*models.User), nil
				}
				return nil, nil
			},
		},
	},
	MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
		c := NewContext(ctx)

		c.EnsureIsAuthorized(nil)

		input := new(struct {
			DisplayName string `validate:"required,min=3,max=50"`
			Picture     string `validate:"url|len=0"`
		})

		c.Validate(inputMap, input, func(e validator.FieldError) string {
			switch e.Field() {
			case "DisplayName":
				return "The displayName field must be between 3 and 50 characters length"
			case "Picture":
				return "The picture field is invalid URL"
			}
			return e.Translate(nil)
		})

		if input.Picture == "" {
			c.User.Picture = null.NewString("", false)
		} else {
			c.User.Picture = null.StringFrom(input.Picture)
		}

		c.User.DisplayName = null.StringFrom(input.DisplayName)
		c.User.UpdatedAt = null.TimeFrom(time.Now())
		c.User.UpdateG(boil.Whitelist("display_name", "picture", "updated_at"))

		return map[string]interface{}{
			"user": c.User,
		}, nil
	},
})

var changePasswordMutation = relay.MutationWithClientMutationID(relay.MutationConfig{
	Name: "ChangePassword",
	InputFields: graphql.InputObjectConfigFieldMap{
		"password": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"currentPassword": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	OutputFields: graphql.Fields{
		"user": &graphql.Field{
			Type: userType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if payload, ok := p.Source.(map[string]interface{}); ok {
					return payload["user"].(*models.User), nil
				}

				return nil, nil
			},
		},
	},
	MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
		c := NewContext(ctx)

		c.EnsureIsAuthorized(nil)

		input := new(struct {
			Password        string `validate:"required,min=8,max=20"`
			CurrentPassword string `validate:"min=8|len=0,max=20"`
		})

		c.Validate(inputMap, input, func(e validator.FieldError) string {
			switch e.Field() {
			case "Password":
				return "The password field must be between 8 and 20 characters length"
			case "CurrentPassword":
				return "The currentPassword field must be between 8 and 20 characters length"
			}
			return e.Translate(nil)
		})

		if c.User.Password.Ptr() != nil && !services.VerifyUserPassword(c.User, input.CurrentPassword) {
			c.Panic(http.StatusBadRequest, "Bad password")
		}

		if err := services.SetUserPassword(c.User, input.Password); err != nil {
			c.Panic(http.StatusInternalServerError, err.Error())
		}

		auth.SetToken(c.EchoCtx, c.User)

		return map[string]interface{}{
			"user": c.User,
		}, nil
	},
})
