package cgo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"time"

	"github.com/volatiletech/null"
	validator "gopkg.in/go-playground/validator.v9"
)

type (
	key int

	// Context for graphql
	Context struct {
		// user
		ID           string      `json:"id" toml:"id" yaml:"id"`
		Email        string      `json:"email" toml:"email" yaml:"email"`
		EmailHash    string      `json:"email_hash" toml:"email_hash" yaml:"email_hash"`
		Username     null.String `json:"username,omitempty" toml:"username" yaml:"username,omitempty"`
		Password     null.String `json:"password,omitempty" toml:"password" yaml:"password,omitempty"`
		DisplayName  null.String `json:"display_name,omitempty" toml:"display_name" yaml:"display_name,omitempty"`
		Picture      null.String `json:"picture,omitempty" toml:"picture" yaml:"picture,omitempty"`
		TokenVersion null.Int64  `json:"token_version,omitempty" toml:"token_version" yaml:"token_version,omitempty"`
		Role         string      `json:"role" toml:"role" yaml:"role"`
		State        string      `json:"state" toml:"state" yaml:"state"`
		CreatedAt    time.Time   `json:"created_at" toml:"created_at" yaml:"created_at"`
		UpdatedAt    null.Time   `json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

		// context
		Request  *http.Request       `json:"-" toml:"-" yaml:"-"`
		Response http.ResponseWriter `json:"-" toml:"-" yaml:"-"`
		Params   map[string]string   `json:"-" toml:"-" yaml:"-"`
		Host     string              `json:"-" toml:"-" yaml:"-"`
	}
)

// Enum values for context_role
const (
	ContextRoleAdmin = "admin"
	ContextRoleStaff = "staff"
	ContextRoleUser  = "user"
)

// Enum values for context_state
const (
	ContextStateEnable      = "enable"
	ContextStateDisable     = "disable"
	ContextStateMaintenance = "maintenance"
)

const (
	contextKeyID           key = 0
	contextKeyEmail        key = 1
	contextKeyEmailHash    key = 2
	contextKeyUsername     key = 3
	contextKeyPassword     key = 4
	contextKeyDisplayName  key = 5
	contextKeyPicture      key = 6
	contextKeyTokenVersion key = 7
	contextKeyRole         key = 8
	contextKeyState        key = 9
	contextKeyCreatedAt    key = 10
	contextKeyUpdatedAt    key = 11
	contextKeyRequest      key = 12
	contextKeyResponse     key = 13
	contextKeyParams       key = 14
	contextKeyHost         key = 15
)

// Validator for validate input
var Validator = validator.New()

// Panic for panic after validation
func (c *Context) Panic(code int, message string) {
	c.Response.WriteHeader(code)
	panic(message)
}

// Param get value of route param by key
func (c *Context) Param(key string) string {
	return c.Params[key]
}

// EnsureIsAuthorized verify if is authorized
func (c *Context) EnsureIsAuthorized(cb func() bool) {
	if c.ID == "" {
		c.Panic(http.StatusUnauthorized, "Anonymous access is denied")
	}

	if c.State == ContextStateMaintenance || (cb != nil && !cb()) {
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

// NewContext create context.Context from Context
func NewContext(c Context) context.Context {
	if c.Request == nil {
		c.Request = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
		c.Request.Header.Set("Content-Type", "application/json")
	}
	if c.Response == nil {
		c.Response = httptest.NewRecorder()
	}
	if c.Params == nil {
		c.Params = map[string]string{}
	}

	r := c.Request
	scheme := "https"
	referer := r.Referer()
	u, err := url.Parse(referer)

	if r.TLS == nil {
		scheme = "http"
	}

	if referer == "" || err != nil {
		c.Host = fmt.Sprintf("%s://%s", scheme, r.Host)
	} else {
		c.Host = fmt.Sprintf("%s://%s", u.Scheme, u.Host)
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, contextKeyID, c.ID)
	ctx = context.WithValue(ctx, contextKeyEmail, c.Email)
	ctx = context.WithValue(ctx, contextKeyEmailHash, c.EmailHash)
	ctx = context.WithValue(ctx, contextKeyUsername, c.Username)
	ctx = context.WithValue(ctx, contextKeyPassword, c.Password)
	ctx = context.WithValue(ctx, contextKeyDisplayName, c.DisplayName)
	ctx = context.WithValue(ctx, contextKeyPicture, c.Picture)
	ctx = context.WithValue(ctx, contextKeyTokenVersion, c.TokenVersion)
	ctx = context.WithValue(ctx, contextKeyRole, c.Role)
	ctx = context.WithValue(ctx, contextKeyState, c.State)
	ctx = context.WithValue(ctx, contextKeyCreatedAt, c.CreatedAt)
	ctx = context.WithValue(ctx, contextKeyUpdatedAt, c.UpdatedAt)
	ctx = context.WithValue(ctx, contextKeyResponse, c.Response)
	ctx = context.WithValue(ctx, contextKeyRequest, c.Request)
	ctx = context.WithValue(ctx, contextKeyHost, c.Host)
	ctx = context.WithValue(ctx, contextKeyParams, c.Params)

	return ctx
}

// FromContext create Context from context.Context
func FromContext(c context.Context) Context {
	return Context{
		ID:           c.Value(contextKeyID).(string),
		Email:        c.Value(contextKeyEmail).(string),
		EmailHash:    c.Value(contextKeyEmailHash).(string),
		Username:     c.Value(contextKeyUsername).(null.String),
		Password:     c.Value(contextKeyPassword).(null.String),
		DisplayName:  c.Value(contextKeyDisplayName).(null.String),
		Picture:      c.Value(contextKeyPicture).(null.String),
		TokenVersion: c.Value(contextKeyTokenVersion).(null.Int64),
		Role:         c.Value(contextKeyRole).(string),
		State:        c.Value(contextKeyState).(string),
		CreatedAt:    c.Value(contextKeyCreatedAt).(time.Time),
		UpdatedAt:    c.Value(contextKeyUpdatedAt).(null.Time),
		Response:     c.Value(contextKeyResponse).(http.ResponseWriter),
		Request:      c.Value(contextKeyRequest).(*http.Request),
		Host:         c.Value(contextKeyHost).(string),
		Params:       c.Value(contextKeyParams).(map[string]string),
	}
}
