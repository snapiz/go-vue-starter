package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/snapiz/go-vue-starter/db/models"

	"github.com/graphql-go/handler"
	"github.com/labstack/echo"
)

// Handler expose graphql
func Handler(c echo.Context) error {
	h := handler.New(&handler.Config{
		Schema:   &Schema,
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
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c = echo.New().NewContext(req, rec)
	}

	if u != nil {
		c.Set("user", u)
	}

	return context.WithValue(context.Background(), contextKeyID, c)
}
