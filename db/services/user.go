package services

import (
	"time"

	"github.com/snapiz/go-vue-starter/common"
	"github.com/snapiz/go-vue-starter/db/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
)

// SetUserPassword set user password
func SetUserPassword(u *models.User, p string) error {
	hash, err := common.Hash(p)

	if err != nil {
		return err
	}

	u.Password = null.StringFrom(hash)
	u.TokenVersion = null.Int64From(time.Now().Unix())
	_, err = u.UpdateG(boil.Whitelist("password", "token_version"))

	return err
}

// VerifyUserPassword verify user password
func VerifyUserPassword(u *models.User, p string) bool {
	ok, err := common.Verify(p, *u.Password.Ptr())

	return ok && err == nil
}
