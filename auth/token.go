package auth

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/snapiz/go-vue-starter/common"
	"github.com/snapiz/go-vue-starter/db/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// SetToken set cookie token
func SetToken(c echo.Context, u models.User) (token string, err error) {
	expireDuration, err := time.ParseDuration(os.Getenv("SESSION_EXPIRE_DURATION"))

	if err != nil {
		return "", err
	}

	token, err = common.SignToken(jwt.StandardClaims{
		Id:      strconv.Itoa(u.ID),
		Subject: strconv.FormatInt(*u.TokenVersion.Ptr(), 10),
		Issuer:  GetIssuer(c),
	})

	if err != nil {
		return "", err
	}

	cookie := &http.Cookie{
		Name:     os.Getenv("SESSION_KEY"),
		Value:    token,
		Expires:  time.Now().Add(expireDuration),
		HttpOnly: true,
		Path:     "/",
	}

	c.SetCookie(cookie)

	return token, nil
}

// RemoveToken unset token from cookie
func RemoveToken(c echo.Context) {
	oldD, _ := time.ParseDuration("-2160h")

	cookie := &http.Cookie{
		Name:     os.Getenv("SESSION_KEY"),
		Value:    "",
		Expires:  time.Now().Add(oldD),
		HttpOnly: true,
		Path:     "/",
	}

	c.SetCookie(cookie)
}

// GetIssuer get issuer from context
func GetIssuer(c echo.Context) string {
	r := c.Request()

	return fmt.Sprintf("%s://%s", c.Scheme(), r.Host)
}
