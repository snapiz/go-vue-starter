package common

import (
	"net/http"

	"github.com/snapiz/go-vue-starter/db/models"

	"github.com/labstack/echo"
	"golang.org/x/net/context"
)

// Context graphql
type Context struct {
	EchoCtx echo.Context
	User    *models.User
}

type key int

// ContextKeyID key
const ContextKeyID key = 0

// EnsureIsAuthorized user is authorized
func (c *Context) EnsureIsAuthorized(cb func(*models.User) bool) {
	if c.User == nil {
		c.EchoCtx.Response().WriteHeader(http.StatusUnauthorized)
		panic("Anonymous access is denied.")
		/* return false */
	}

	if cb != nil && !cb(c.User) {
		c.EchoCtx.Response().WriteHeader(http.StatusForbidden)
		panic("Access is denied.")
	}
}

// NewContext create graphql context
func NewContext(c context.Context) (ctx Context) {
	ctx = Context{
		EchoCtx: c.Value(ContextKeyID).(echo.Context),
	}

	if user := ctx.EchoCtx.Get("user"); user != nil {
		ctx.User = ctx.EchoCtx.Get("user").(*models.User)
	}

	return ctx
}
