package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"

	"github.com/snapiz/go-vue-starter/services/main/src/utils"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/snapiz/go-vue-starter/services/main/src/models"
	"github.com/spf13/cobra"
)

func init() {
	var email string
	var username string
	var password string
	var role string

	c := &cobra.Command{
		Use:   "user:create",
		Short: "Create user",
		Run: func(cmd *cobra.Command, args []string) {
			hash := md5.Sum([]byte(email))
			u := &models.User{
				Email:     email,
				EmailHash: hex.EncodeToString(hash[:]),
				Username:  null.StringFrom(username),
				Role:      role,
			}
			utils.SetUserPassword(u, password)
			err := u.InsertG(boil.Whitelist("email", "email_hash", "username", "password", "role", "token_version"))

			if err != nil {
				log.Fatal(err)
			}
		},
	}

	c.PersistentFlags().StringVarP(&email, "email", "e", "", "User email")
	c.PersistentFlags().StringVarP(&username, "username", "u", "", "Username")
	c.PersistentFlags().StringVarP(&password, "password", "p", "azertyui", "User password")
	c.PersistentFlags().StringVarP(&role, "role", "r", "user", "Default is user")

	root.AddCommand(c)
}
