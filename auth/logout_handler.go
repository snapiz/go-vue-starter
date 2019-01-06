package auth

import (
	"net/http"

	"github.com/labstack/echo"
)

// LogoutHandler auth logout
var LogoutHandler = func(c echo.Context) error {
	RemoveToken(c)

	return c.JSON(http.StatusOK, "ok")
}
