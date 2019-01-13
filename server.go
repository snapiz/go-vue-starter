package main

import (
	"os"

	_ "github.com/bmizerany/pq"
	"github.com/labstack/echo"

	"github.com/snapiz/go-vue-starter/server/api"
	"github.com/snapiz/go-vue-starter/server/auth"
	_ "github.com/snapiz/go-vue-starter/server/db"
	"github.com/snapiz/go-vue-starter/server/middlewares"
)

func main() {
	e := echo.New()

	e.Any("/auth/logout", auth.LogoutHandler)
	e.Any("/auth/local", auth.LocalHandler)
	e.Any("/auth/local/register", auth.RegisterHandler)
	e.GET("/auth/:provider", auth.OAuth2Handler)
	e.GET("/auth/:provider/callback", auth.OAuth2CallbackHandler)

	e.Any("/graphql", api.Handler, middlewares.JWT)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
