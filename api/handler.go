package api

import (
	"context"
	"net/http"

	"github.com/snapiz/go-vue-starter/api/common"
	"github.com/snapiz/go-vue-starter/db/models"

	"github.com/snapiz/go-vue-starter/api/schema"

	"github.com/graphql-go/handler"
	"github.com/labstack/echo"
)

// Handler expose graphql
func Handler(c echo.Context) error {
	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	f := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := NewGraphQLContext(c, nil)

		h.ContextHandler(ctx, w, r)
	})

	return echo.WrapHandler(f)(c)
}

// NewGraphQLContext graphql context
func NewGraphQLContext(c echo.Context, u *models.User) (ctx context.Context) {
	if c == nil {
		c = echo.New().NewContext(nil, nil)
	}

	if u != nil {
		c.Set("user", u)
	}

	return context.WithValue(context.Background(), common.ContextKeyID, c)
}
