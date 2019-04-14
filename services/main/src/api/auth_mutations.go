package api

import (
	"context"
	"net/http"

	"github.com/snapiz/go-vue-starter/packages/cgo"
	"github.com/snapiz/go-vue-starter/services/main/src/db"
	"github.com/snapiz/go-vue-starter/services/main/src/models"
	"github.com/snapiz/go-vue-starter/services/main/src/utils"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/oauth2"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	validator "gopkg.in/go-playground/validator.v9"
)

var loginWithOAuth2Mutation = relay.MutationWithClientMutationID(relay.MutationConfig{
	Name: "LoginWithOAuth2",
	InputFields: graphql.InputObjectConfigFieldMap{
		"provider": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(oAuth2providerType),
		},
		"code": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	OutputFields: graphql.Fields{
		"user": &graphql.Field{
			Type: userType,
		},
	},
	MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
		c := cgo.FromContext(ctx)

		input := new(struct {
			Code     string `validate:"required"`
			Provider string `validate:"required"`
		})

		c.Validate(inputMap, input, func(e validator.FieldError) string {
			return "errors.auth.missingFields"
		})

		conf, provider := utils.CreateOauthConf(&c, input.Provider)
		tok, err := conf.Exchange(oauth2.NoContext, input.Code)

		if err != nil {
			c.Panic(http.StatusInternalServerError, err.Error())
		}

		p, err := provider.GetProfile(tok)

		if err != nil {
			c.Panic(http.StatusInternalServerError, "Failed to get profile")
		}

		users, err := models.Users(qm.Where("email = ?", p.Email)).AllG()

		if err != nil {
			c.Panic(http.StatusInternalServerError, "Failed to fetch user")
		}

		var user *models.User

		if users == nil {
			user = &models.User{
				Email:       p.Email,
				DisplayName: null.StringFrom(p.DisplayName),
				Picture:     null.StringFrom(p.Picture),
			}

			err := db.CreateUser(user)

			if err != nil {
				c.Panic(http.StatusInternalServerError, "Failed to create user")
			}
		} else {
			user = users[0]
		}

		// Check account is disabled
		if user.State == models.UserStateDisable {
			c.Panic(http.StatusUnauthorized, "errors.auth.accountIsDisabled")
		}

		userProviders, err := models.UserProviders(qm.Where("provider = ? AND provider_id = ? AND user_id = ?", input.Provider, p.ID, user.ID)).AllG()

		if err != nil {
			c.Panic(http.StatusInternalServerError, "Failed to fetch user provider")
		}

		if userProviders == nil {
			userProvider := &models.UserProvider{
				Provider:   input.Provider,
				ProviderID: p.ID,
				UserID:     user.ID,
			}

			err = userProvider.InsertG(boil.Whitelist("provider", "provider_id", "user_id"))

			if err != nil {
				c.Panic(http.StatusInternalServerError, "Failed to create user provider")
			}
		}

		_, err = c.SetUser(user)

		if err != nil {
			c.Panic(http.StatusInternalServerError, "Failed to set token")
		}

		return map[string]interface{}{
			"user": c,
		}, nil
	},
})
