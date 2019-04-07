package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/snapiz/go-vue-starter/packages/cgo"
	_ "github.com/snapiz/go-vue-starter/services/main/src/db"
	"github.com/snapiz/go-vue-starter/services/main/src/models"
	"github.com/snapiz/go-vue-starter/services/main/src/schema"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func main() {
	cgo.Start(cgo.ServiceConfig{
		Name:   "main",
		Schema: schema.Schema,
		FetchUser: func(queryMod qm.QueryMod) (interface{}, error) {
			users, err := models.Users(queryMod).AllG()

			if err != nil {
				return nil, err
			}

			if users == nil {
				return nil, nil
			}

			return users[0], nil
		},
	}, func(r *mux.Router) {
		r.HandleFunc("/_login/{username}", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			u, err := models.Users(qm.Where("username = ?", vars["username"])).OneG()

			if err != nil {
				email := vars["username"] + "@gvs.com"
				hash := md5.Sum([]byte(email))
				u = &models.User{
					Email:        email,
					EmailHash:    hex.EncodeToString(hash[:]),
					TokenVersion: null.Int64From(time.Now().Unix()),
					DisplayName:  null.StringFrom("John doe"),
					Username:     null.StringFrom(vars["username"]),
				}
				u.InsertG(boil.Whitelist("email", "email_hash", "token_version", "display_name", "username"))
			}

			c := cgo.Context{
				ID:           u.ID,
				Email:        u.Email,
				EmailHash:    u.EmailHash,
				Username:     u.Username,
				Password:     u.Password,
				TokenVersion: u.TokenVersion,
				Picture:      u.Picture,
				DisplayName:  u.DisplayName,
				State:        u.State,
				Role:         u.Role,
				CreatedAt:    u.CreatedAt,
				UpdatedAt:    u.UpdatedAt,
				Request:      r,
				Response:     w,
			}
			c.Host = "http://localhost:9000"
			c.CreateToken()

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "ID: %v\n", u.ID)
		})
	})
}
