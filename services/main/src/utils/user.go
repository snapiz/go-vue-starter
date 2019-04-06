package utils

import (
	"time"

	"github.com/snapiz/go-vue-starter/packages/cgo"
	"github.com/snapiz/go-vue-starter/services/main/src/models"
	"github.com/volatiletech/null"
)

// SetUserPassword set user password
func SetUserPassword(u *models.User, p string) error {
	hash, err := cgo.Hash(p)

	if err != nil {
		return err
	}

	u.Password = null.StringFrom(hash)
	u.TokenVersion = null.Int64From(time.Now().Unix())

	return err
}

// VerifyUserPassword verify user password
func VerifyUserPassword(hash string, p string) bool {
	ok, err := cgo.Verify(p, hash)

	return ok && err == nil
}
