package db

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/snapiz/go-vue-starter/packages/cgo"
	"github.com/snapiz/go-vue-starter/services/main/src/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
)

// HashUserPassword hash plain password
func HashUserPassword(user *models.User) error {
	hash, err := cgo.Hash(*user.Password.Ptr())

	if err != nil {
		return err
	}

	user.Password = null.StringFrom(hash)
	user.TokenVersion = null.Int64From(time.Now().Unix())

	return err
}

// VerifyUserPassword verify user password
func VerifyUserPassword(hash string, p string) bool {
	ok, err := cgo.Verify(p, hash)

	return ok && err == nil
}

// CreateUser with minimal requirements
func CreateUser(user *models.User) error {
	hash := md5.Sum([]byte(user.Email))
	user.EmailHash = hex.EncodeToString(hash[:])

	if err := HashUserPassword(user); err != nil {
		return err
	}

	return user.InsertG(boil.Whitelist("email", "email_hash", "username", "password", "role", "token_version"))
}
