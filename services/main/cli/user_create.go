package main

import (
	"log"

	"github.com/volatiletech/null"

	"github.com/snapiz/go-vue-starter/services/main/src/db"
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
			user := &models.User{
				Email:    email,
				Role:     role,
				Password: null.StringFrom(password),
			}

			if username != "" {
				user.Username = null.StringFrom(username)
			}

			if err := db.CreateUser(user); err != nil {
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
