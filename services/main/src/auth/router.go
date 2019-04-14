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
}
