package auth

import (
	"github.com/snapiz/go-vue-starter/packages/cgo"
)

// AddRoutes to cgo.Router
func AddRoutes(r *cgo.Router) {
	r.Add(&cgo.RouteConfig{
		Path:        "/auth/local",
		HandlerFunc: LocalHandler,
	})

	r.Add(&cgo.RouteConfig{
		Path:        "/auth/local/new",
		HandlerFunc: LocalNewHandler,
	})

	r.Add(&cgo.RouteConfig{
		Path:        "/auth/logout",
		HandlerFunc: LogoutHandler,
	})

	r.Add(&cgo.RouteConfig{
		Path:        "/auth/{provider:google|facebook}",
		HandlerFunc: OAuth2Handler,
	})

	r.Add(&cgo.RouteConfig{
		Path:        "/auth/{provider:google|facebook}/callback",
		HandlerFunc: OAuth2CallbackHandler,
	})
}
