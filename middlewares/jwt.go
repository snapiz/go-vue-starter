package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"github.com/snapiz/go-vue-starter/auth"
	"github.com/snapiz/go-vue-starter/common"
	"github.com/snapiz/go-vue-starter/db/models"

	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/labstack/echo"
)

// JWT middleware verify auth token.
func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authCookie, err := c.Cookie(os.Getenv("SESSION_KEY"))

		if err != nil {
			return next(c)
		}

		claims, err := common.VerifyToken(authCookie.Value, auth.GetIssuer(c))

		if err != nil {
			return next(c)
		}

		users, err := models.Users(qm.Where("id = ? AND token_version = ? AND role != 'disable'", claims.Id, claims.Subject)).AllG()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintln(err))
		}

		if users == nil {
			auth.RemoveToken(c)
		} else {
			u := users[0]
			auth.SetToken(c, *u)
			c.Set("user", u)
		}

		return next(c)
	}
}
