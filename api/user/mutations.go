package user

import (
	"github.com/snapiz/go-vue-starter/auth"
	"context"
	"net/http"
	"time"

	"github.com/snapiz/go-vue-starter/db/models"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"github.com/snapiz/go-vue-starter/api/common"
	com "github.com/snapiz/go-vue-starter/common"
	validator "gopkg.in/go-playground/validator.v9"
)

// UpdateUserMutation update current user
var UpdateUserMutation = relay.MutationWithClientMutationID(relay.MutationConfig{
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
			Type: UserType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if payload, ok := p.Source.(map[string]interface{}); ok {
					return payload["user"].(*models.User), nil
				}
				return nil, nil
			},
		},
	},
	MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
		c := common.NewContext(ctx)

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

// ChangePasswordMutation change current user password
var ChangePasswordMutation = relay.MutationWithClientMutationID(relay.MutationConfig{
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
			Type: UserType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if payload, ok := p.Source.(map[string]interface{}); ok {
					return payload["user"].(*models.User), nil
				}

				return nil, nil
			},
		},
	},
	MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
		c := common.NewContext(ctx)

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

		if c.User.Password.Ptr() != nil {
			if ok, err := com.Verify(input.CurrentPassword, *c.User.Password.Ptr()); !ok || err != nil {
				c.Panic(http.StatusUnauthorized, "Bad password")
			}
		}

		hash, err := com.Hash(input.Password)

		if err != nil {
			c.Panic(http.StatusInternalServerError, err.Error())
		}

		c.User.Password = null.StringFrom(hash)
		c.User.TokenVersion = null.Int64From(time.Now().Unix())
		_, err = c.User.UpdateG(boil.Whitelist("password", "token_version"))

		if err != nil {
			c.Panic(http.StatusInternalServerError, err.Error())
		}

		_, err = auth.SetToken(c.EchoCtx, *c.User)

		return map[string]interface{}{
			"user": c.User,
		}, nil
	},
})
