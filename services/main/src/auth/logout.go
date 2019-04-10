package auth

import (
	"github.com/snapiz/go-vue-starter/packages/cgo"
)

// LogoutHandler for local auth
func LogoutHandler(context cgo.Context) (map[string]interface{}, error) {
	context.RemoveToken()

	return map[string]interface{}{
		"success": true,
	}, nil
}
