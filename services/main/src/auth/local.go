package auth

import (
	"net/http"

	"github.com/snapiz/go-vue-starter/packages/cgo"
	"github.com/snapiz/go-vue-starter/services/main/src/db"
	"github.com/snapiz/go-vue-starter/services/main/src/models"
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
		return "errors.auth.missingFields"
	})

	users, err := models.Users(qm.Where("email = ?", input.Login), qm.Or("username = ?", input.Login)).AllG()

	if err != nil {
		context.Panic(http.StatusInternalServerError, "Falied to request user")
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
		"me": context,
	}, nil
}
