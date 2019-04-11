package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/snapiz/go-vue-starter/packages/cgo"
	"github.com/snapiz/go-vue-starter/services/main/src/db"
	"github.com/snapiz/go-vue-starter/services/main/src/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/queries/qm"
	validator "gopkg.in/go-playground/validator.v9"
)

// LocalHandler for local auth
func LocalHandler(context cgo.Context) (map[string]interface{}, error) {
	input := new(struct {
		Login    string `json:"login" validate:"required"`
		Password string `json:"password" validate:"required"`
	})

	context.Validate(nil, input, func(e validator.FieldError) string {
		return fmt.Sprintf("The %s field is required", strings.ToLower(e.Field()))
	})

	users, err := models.Users(qm.Where("email = ? or username = ?", input.Login, input.Login)).AllG()

	if err != nil {
		context.Panic(http.StatusInternalServerError, "Failed to request user")
	}

	if users == nil {
		context.Panic(http.StatusUnauthorized, "errors.auth.invalidCredentials")
	}

	user := users[0]

	if user.Password.Ptr() == nil {
		context.Panic(http.StatusUnauthorized, "errors.auth.invalidCredentials")
	}

	// Verify password
	if ok := db.VerifyUserPassword(*user.Password.Ptr(), input.Password); !ok {
		context.Panic(http.StatusUnauthorized, "errors.auth.invalidCredentials")
	}

	// Check account is disabled
	if user.State == models.UserStateDisable {
		context.Panic(http.StatusUnauthorized, "errors.auth.accountIsDisabled")
	}

	_, err = context.SetUser(user)

	if err != nil {
		context.Panic(http.StatusInternalServerError, "Failed to write token")
	}

	return map[string]interface{}{
		"id": context.ID,
		"role": context.Role,
	}, nil
}

// LocalNewHandler for sign up
func LocalNewHandler(context cgo.Context) (map[string]interface{}, error) {
	input := new(struct {
		Email    string `json:"email" validate:"required,email"`
		Username string `json:"username" validate:"required,alphanum,min=3,max=50"`
		Password string `json:"password" validate:"required,min=8,max=20"`
	})

	context.Validate(nil, input, func(e validator.FieldError) string {
		switch e.Field() {
		case "Email":
			return "The email field must be an email"
		case "Username":
			return "The username field must be between 3 and 50 characters long"
		case "Password":
			return "The password field must be between 2 and 20 characters long"
		default:
			return e.Translate(nil)
		}
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
		"id": context.ID,
		"role": context.Role,
	}, nil
}
