package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/snapiz/go-vue-starter/packages/cgo"
	"github.com/snapiz/go-vue-starter/services/main/src/db"
	"github.com/snapiz/go-vue-starter/services/main/src/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	validator "gopkg.in/go-playground/validator.v9"
)

type profile struct {
	ID          string
	DisplayName string
	Picture     string
	Email       string
}

type provider struct {
	ClientID     string
	ClientSecret string
	Endpoint     oauth2.Endpoint
	Scopes       []string
	GetProfile   func(*oauth2.Token) (*profile, error)
}

var providers = map[string]*provider{
	"google": &provider{
		ClientID:     "GOOGLE_CLIENT_ID",
		ClientSecret: "GOOGLE_CLIENT_SECRET",
		Endpoint:     google.Endpoint,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		GetProfile: func(tok *oauth2.Token) (*profile, error) {
			response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + tok.AccessToken)
			if err != nil {
				return nil, fmt.Errorf("failed getting user info: %s", err.Error())
			}
			defer response.Body.Close()

			p := new(struct {
				ID      string `json:"id"`
				Name    string `json:"name"`
				Picture string `json:"picture"`
				Email   string `json:"email"`
			})
			json.NewDecoder(response.Body).Decode(&p)

			return &profile{
				ID:          p.ID,
				Email:       p.Email,
				DisplayName: p.Name,
				Picture:     p.Picture,
			}, nil
		},
	},
	"facebook": &provider{
		ClientID:     "FACEBOOK_CLIENT_ID",
		ClientSecret: "FACEBOOK_CLIENT_SECRET",
		Endpoint:     facebook.Endpoint,
		Scopes:       []string{"public_profile", "email"},
		GetProfile: func(tok *oauth2.Token) (*profile, error) {
			response, err := http.Get("https://graph.facebook.com/me?fields=id,email,name,picture&access_token=" + tok.AccessToken)
			if err != nil {
				return nil, fmt.Errorf("failed getting user info: %s", err.Error())
			}
			defer response.Body.Close()

			type facebookPictureData struct {
				URL string `json:"url"`
			}

			type facebookPicture struct {
				Data facebookPictureData `json:"data"`
			}

			p := new(struct {
				ID      string          `json:"id"`
				Name    string          `json:"name"`
				Picture facebookPicture `json:"picture"`
				Email   string          `json:"email"`
			})

			json.NewDecoder(response.Body).Decode(&p)

			return &profile{
				ID:          p.ID,
				Email:       p.Email,
				DisplayName: p.Name,
				Picture:     p.Picture.Data.URL,
			}, nil
		},
	},
}

func createOauthConf(c *cgo.Context, p *provider) oauth2.Config {
	name := c.Param("provider")

	return oauth2.Config{
		ClientID:     os.Getenv(p.ClientID),
		ClientSecret: os.Getenv(p.ClientSecret),
		RedirectURL:  fmt.Sprintf("%s/auth/%s/callback", c.Host, name),
		Scopes:       p.Scopes,
		Endpoint:     p.Endpoint,
	}
}

// OAuth2Handler for oauth2 auth
func OAuth2Handler(context cgo.Context) (map[string]interface{}, error) {
	name := context.Param("provider")
	provider := providers[name]

	if provider == nil {
		context.Panic(http.StatusBadRequest, fmt.Sprintf("Unknown provider %s.", name))
	}

	conf := createOauthConf(&context, provider)

	return context.Redirect(conf.AuthCodeURL("state"))
}

// OAuth2CallbackHandler for oauth2 callback
func OAuth2CallbackHandler(context cgo.Context) (map[string]interface{}, error) {
	input := new(struct {
		Code        string `form:"code"`
		ClientID    string `form:"clientId"`
		RedirectURI string `form:"redirectUri"`
	})

	context.Validate(nil, input, func(e validator.FieldError) string {
		return "errors.auth.missingFields"
	})

	name := context.Param("provider")
	provider := providers[name]

	if provider == nil {
		context.Panic(http.StatusBadRequest, fmt.Sprintf("Unknown provider %s.", name))
	}

	conf := createOauthConf(&context, provider)
	tok, err := conf.Exchange(oauth2.NoContext, input.Code)

	if err != nil {
		context.Panic(http.StatusInternalServerError, "Failed to exchange")
	}

	p, err := provider.GetProfile(tok)

	if err != nil {
		context.Panic(http.StatusInternalServerError, "Failed to get profile")
	}

	users, err := models.Users(qm.Where("email = ?", p.Email)).AllG()

	if err != nil {
		context.Panic(http.StatusInternalServerError, "Failed to fetch user")
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
			context.Panic(http.StatusInternalServerError, "Failed to create user")
		}
	} else {
		user = users[0]
	}

	// Check account is disabled
	if user.State == models.UserStateDisable {
		context.Panic(http.StatusUnauthorized, "errors.auth.accountIsDisabled")
	}

	userProviders, err := models.UserProviders(qm.Where("provider = ? AND provider_id = ? AND user_id = ?", name, p.ID, user.ID)).AllG()

	if err != nil {
		context.Panic(http.StatusInternalServerError, "Failed to fetch user provider")
	}

	if userProviders == nil {
		userProvider := &models.UserProvider{
			Provider:   name,
			ProviderID: p.ID,
			UserID:     user.ID,
		}

		err = userProvider.InsertG(boil.Whitelist("provider", "provider_id", "user_id"))

		if err != nil {
			context.Panic(http.StatusInternalServerError, "Failed to create user provider")
		}
	}

	_, err = context.SetUser(user)

	if err != nil {
		context.Panic(http.StatusInternalServerError, "Failed to set token")
	}

	return context.Redirect(context.Host)
}
