package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/snapiz/go-vue-starter/packages/cgo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
)

type profile struct {
	ID          string
	DisplayName string
	Picture     string
	Email       string
}

type Provider struct {
	ClientID     string
	ClientSecret string
	Endpoint     oauth2.Endpoint
	Scopes       []string
	GetProfile   func(*oauth2.Token) (*profile, error)
}

var providers = map[string]*Provider{
	"google": &Provider{
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
	"facebook": &Provider{
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

// CreateOauthConf create oauth2 config
func CreateOauthConf(c *cgo.Context, name string) (*oauth2.Config, *Provider) {
	p := providers[name]

	return &oauth2.Config{
		ClientID:     os.Getenv(p.ClientID),
		ClientSecret: os.Getenv(p.ClientSecret),
		RedirectURL:  fmt.Sprintf("%s/oauth2/%s", c.Origin, name),
		Scopes:       p.Scopes,
		Endpoint:     p.Endpoint,
	}, p
}

// GetProviderAuthURLs return mapped provider auth url
func GetProviderAuthURLs(c *cgo.Context) map[string]string  {
	urls := map[string]string{};

	for k := range providers {
		conf, _ := CreateOauthConf(c, k)
		urls[k] = conf.AuthCodeURL("state")
	}

	return urls
}