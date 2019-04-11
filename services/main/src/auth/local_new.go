package auth

import (
	"net/http"

	"github.com/snapiz/go-vue-starter/packages/cgo"
	"github.com/snapiz/go-vue-starter/services/main/src/db"
	"github.com/snapiz/go-vue-starter/services/main/src/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/queries/qm"
	validator "gopkg.in/go-playground/validator.v9"
)

// LocalNewHandler for sign up
func LocalNewHandler(context cgo.Context) (map[string]interface{}, error) {
	input := new(struct {
		Email    string `json:"email" validate:"required,email"`
		Username string `json:"username" validate:"required,alphanum,min=3,max=50"`
		Password string `json:"password" validate:"required,min=8,max=20"`
	})

	context.Validate(nil, input, func(e validator.FieldError) string {
		return "errors.auth.missingFields"
	})

	users, err := models.Users(qm.Where("email = ? or username = ?", input.Email, input.Username)).AllG()

	if err != nil {
		context.Panic(http.StatusInternalServerError, "Failed to fetch user by email or username")
	}

	if users != nil {
		if users[0].Email == input.Email {
			context.Panic(http.StatusBadRequest, "errors.auth.emailAlreadyExists")
		} else {
			context.Panic(http.StatusBadRequest, "errors.auth.usernameAlreadyExists")
		}
	}

	user := &models.User{
		Email:    input.Email,
		Username: null.StringFrom(input.Username),
		Role:     "user",
		Password: null.StringFrom(input.Password),
	}

	if err := db.CreateUser(user); err != nil {
		context.Panic(http.StatusInternalServerError, "Failed to create user")
	}

	_, err = context.SetUser(user)

	if err != nil {
		context.Panic(http.StatusInternalServerError, "Failed to write token")
	}

	return map[string]interface{}{
		"me": context,
	}, nil
}
