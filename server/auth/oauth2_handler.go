package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/snapiz/go-vue-starter/server/db/models"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/labstack/echo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
)

type oauth2InputData struct {
	Code        string `json:"code"`
	ClientID    string `json:"clientId"`
	RedirectURI string `json:"redirectUri"`
}

type profile struct {
	ID          string
	DisplayName string
	Picture     string
	Email       string
}

type providerConfig struct {
	ClientID     string
	ClientSecret string
	Endpoint     oauth2.Endpoint
	Scopes       []string
	GetProfile   func(*oauth2.Token) (*profile, error)
}

var providers = map[string]*providerConfig{
	"google": &providerConfig{
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

			type googlePlusProfile struct {
				ID      string `json:"id"`
				Name    string `json:"name"`
				Picture string `json:"picture"`
				Email   string `json:"email"`
			}

			gPP := &googlePlusProfile{}
			json.NewDecoder(response.Body).Decode(&gPP)

			return &profile{
				ID:          gPP.ID,
				Email:       gPP.Email,
				DisplayName: gPP.Name,
				Picture:     gPP.Picture,
			}, nil
		},
	},
	"facebook": &providerConfig{
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

			type facebookProfile struct {
				ID      string          `json:"id"`
				Name    string          `json:"name"`
				Picture facebookPicture `json:"picture"`
				Email   string          `json:"email"`
			}

			gPP := &facebookProfile{}
			json.NewDecoder(response.Body).Decode(&gPP)

			return &profile{
				ID:          gPP.ID,
				Email:       gPP.Email,
				DisplayName: gPP.Name,
				Picture:     gPP.Picture.Data.URL,
			}, nil
		},
	},
}

// OAuth2Handler redirect to permission
func OAuth2Handler(c echo.Context) error {
	name := c.Param("provider")
	provider := providers[name]

	if provider == nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Unknown provider %s.", name))
	}

	conf := createOauthConf(c, provider)

	return c.Redirect(http.StatusTemporaryRedirect, conf.AuthCodeURL("state"))
}

// OAuth2CallbackHandler oauth2 callback
func OAuth2CallbackHandler(c echo.Context) error {
	input := &oauth2InputData{}

	if err := c.Bind(input); err != nil {
		return c.String(http.StatusBadRequest, "Invalid input")
	}

	name := c.Param("provider")
	provider := providers[name]

	if provider == nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Unknown provider %s.", name))
	}

	conf := createOauthConf(c, provider)
	tok, err := conf.Exchange(oauth2.NoContext, input.Code)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("%s.", err.Error()))
	}

	p, err := provider.GetProfile(tok)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("%s.", err.Error()))
	}

	users, err := models.Users(qm.Where("email = ?", p.Email)).AllG()

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("%s.", err.Error()))
	}

	var user *models.User

	if users == nil {
		user = &models.User{
			Email:        p.Email,
			TokenVersion: null.Int64From(time.Now().Unix()),
			DisplayName:  null.StringFrom(p.DisplayName),
			Picture:      null.StringFrom(p.Picture),
		}

		err := user.InsertG(boil.Whitelist("email", "token_version", "display_name", "picture"))

		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("%s.", err.Error()))
		}
	} else {
		user = users[0]
	}

	userProviders, err := models.UserProviders(qm.Where("provider = ? AND provider_id = ? AND user_id = ?", name, p.ID, user.ID)).AllG()

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("%s.", err.Error()))
	}

	if userProviders == nil {
		userProvider := &models.UserProvider{
			Provider:   name,
			ProviderID: p.ID,
			UserID:     user.ID,
		}

		err = userProvider.InsertG(boil.Whitelist("provider", "provider_id", "user_id"))

		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("%s.", err.Error()))
		}
	}

	_, err = SetToken(c, user)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("%s.", err.Error()))
	}

	return c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s", getHost(c)))
}

func createOauthConf(c echo.Context, provider *providerConfig) oauth2.Config {
	name := c.Param("provider")

	return oauth2.Config{
		ClientID:     os.Getenv(provider.ClientID),
		ClientSecret: os.Getenv(provider.ClientSecret),
		RedirectURL:  fmt.Sprintf("%s/auth/%s/callback", getHost(c), name),
		Scopes:       provider.Scopes,
		Endpoint:     provider.Endpoint,
	}
}

func getHost(c echo.Context) string {
	r := c.Request()
	u, err := url.Parse(r.Referer())

	if err != nil {
		return GetIssuer(c)
	}

	return fmt.Sprintf("%s://%s", u.Scheme, u.Host)
}
