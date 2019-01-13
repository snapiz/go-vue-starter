package api

import (
	"encoding/json"
	"net/http"

	"github.com/snapiz/go-vue-starter/server/db/models"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/labstack/echo"
	"golang.org/x/net/context"
)

// Context graphql
type Context struct {
	EchoCtx  echo.Context
	User     *models.User
}

type key int

const contextKeyID key = 0

// Validator instance
var Validator = validator.New()

// Panic error
func (c *Context) Panic(code int, message string) {
	c.EchoCtx.Response().WriteHeader(code)
	panic(message)
}

// EnsureIsAuthorized user is authorized
func (c *Context) EnsureIsAuthorized(cb func(*models.User) bool) {
	if c.User == nil {
		c.Panic(http.StatusUnauthorized, "Anonymous access is denied")
	}

	if cb != nil && !cb(c.User) {
		c.Panic(http.StatusForbidden, "Access is denied")
	}
}

// Validate struct
func (c *Context) Validate(inputMap map[string]interface{}, s interface{}, cb func(err validator.FieldError) string) {
	jsonString, _ := json.Marshal(inputMap)
	json.Unmarshal(jsonString, &s)

	if err := Validator.Struct(s); err != nil {
		e := err.(validator.ValidationErrors)[0].(validator.FieldError)

		if cb == nil {
			c.Panic(http.StatusBadRequest, e.Translate(nil))
		} else {
			c.Panic(http.StatusBadRequest, cb(e))
		}
	}
}

// NewContext create graphql context
func NewContext(c context.Context) (ctx Context) {
	ctx = Context{
		EchoCtx: c.Value(contextKeyID).(echo.Context),
	}

	if user := ctx.EchoCtx.Get("user"); user != nil {
		ctx.User = ctx.EchoCtx.Get("user").(*models.User)
	}

	return ctx
}
