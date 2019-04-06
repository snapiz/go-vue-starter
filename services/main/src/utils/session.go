package utils

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/snapiz/go-vue-starter/packages/cgo"
	"github.com/volatiletech/sqlboiler/queries/qm"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/snapiz/go-vue-starter/services/main/src/models"
)

const (
	sessionExpireDuraction = "SESSION_EXPIRE_DURATION"
	sessionKey             = "SESSION_KEY"
)

// SetToken set cookie token
func SetToken(c cgo.Context) (token string, err error) {
	expireDuration, err := time.ParseDuration(os.Getenv(sessionExpireDuraction))

	if err != nil {
		return "", err
	}

	token, err = cgo.SignToken(jwt.StandardClaims{
		Id:      c.ID,
		Subject: strconv.FormatInt(*c.TokenVersion.Ptr(), 10),
		Issuer:  c.Host,
	})

	if err != nil {
		return "", err
	}

	cookie := &http.Cookie{
		Name:     os.Getenv(sessionKey),
		Value:    token,
		Expires:  time.Now().Add(expireDuration),
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(c.Response, cookie)

	return token, nil
}

// RemoveToken unset token from cookie
func RemoveToken(c cgo.Context) {
	cookie := &http.Cookie{
		Name:     os.Getenv(sessionKey),
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Path:     "/",
	}

	c.Response.Header().Del("Set-Cookie")
	http.SetCookie(c.Response, cookie)
}

// NewContextWithUser bind user to context
func NewContextWithUser(c cgo.Context) cgo.Context {
	authCookie, err := c.Request.Cookie(os.Getenv(sessionKey))

	if err != nil {
		return c
	}

	claims, err := cgo.VerifyToken(authCookie.Value, c.Host)

	if err != nil {
		c.Panic(http.StatusInternalServerError, "failed to verify token")
	}

	users, err := models.Users(qm.Where("id = ? AND token_version = ? AND state != 'disable'", claims.Id, claims.Subject)).AllG()

	if err != nil {
		c.Panic(http.StatusInternalServerError, "failed to fetch user")
	}

	if users == nil {
		RemoveToken(c)
		return c
	}

	u := users[0]

	c.ID = u.ID
	c.Email = u.Email
	c.EmailHash = u.EmailHash
	c.Username = u.Username
	c.Password = u.Password
	c.TokenVersion = u.TokenVersion
	c.Picture = u.Picture
	c.DisplayName = u.DisplayName
	c.State = u.State
	c.Role = u.Role
	c.CreatedAt = u.CreatedAt
	c.UpdatedAt = u.UpdatedAt

	SetToken(c)

	return c
}
