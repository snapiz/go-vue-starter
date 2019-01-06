package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/snapiz/go-vue-starter/common"
	"github.com/snapiz/go-vue-starter/db/models"

	"github.com/labstack/echo"
)

type respJSON struct {
	ID int `json:"id"`
	Email string `json:"email"`
	DisplayName null.String `json:"display_name"`
	Picture null.String `json:"picture"`
	Role string `json:"role"`
}

type localParams struct {
	Email    string `json:"email" form:"email" query:"email" validate:"required,email"`
	Password string `json:"password" form:"password" query:"password" validate:"required,min=8"`
}

// LocalHandler auth local
var LocalHandler = func(c echo.Context) error {
	p := new(localParams)

	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// validate params
	if err := common.Validate(p); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	users, err := models.Users(qm.Where("email = ?", p.Email)).AllG()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintln(err))
	}

	if users == nil {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials.")
	}

	user := *users[0]

	if user.Password.Ptr() == nil {
		return c.JSON(http.StatusUnauthorized, "Password not defined.")
	}

	// Verify password
	if ok, err := common.Verify(p.Password, *user.Password.Ptr()); !ok || err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials.")
	}

	// Check account is disabled
	if user.Role == models.UserRoleDisable {
		return c.JSON(http.StatusUnauthorized, "Account disabled.")
	}

	// write token into cookie
	_, err = SetToken(c, user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintln(err))
	}

	return c.JSON(http.StatusOK, &respJSON{
		ID: user.ID,
		Email: user.Email,
		DisplayName: user.DisplayName,
		Picture: user.Picture,
		Role: user.Role,
	})
}

type registerParams struct {
	Email    string `json:"email" form:"email" query:"email" validate:"required,email"`
	Password string `json:"password" form:"password" query:"password" validate:"required,min=8,max=20"`
}

// RegisterHandler auth local
var RegisterHandler = func(c echo.Context) error {
	p := new(registerParams)

	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := common.Validate(p); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	users, err := models.Users(qm.Where("email = ?", p.Email)).AllG()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintln(err))
	}

	if users != nil {
		return c.JSON(http.StatusBadRequest, "Email already exists.")
	}

	hash, err := common.Hash(p.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintln(err))
	}

	user := models.User{
		Email:        p.Email,
		Password:     null.StringFrom(hash),
		TokenVersion: null.Int64From(time.Now().Unix()),
	}

	err = user.InsertG(boil.Whitelist("email", "password", "token_version"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintln(err))
	}

	_, err = SetToken(c, user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintln(err))
	}

	return c.JSON(http.StatusOK, &respJSON{
		ID: user.ID,
		Email: user.Email,
		DisplayName: user.DisplayName,
		Picture: user.Picture,
		Role: user.Role,
	})
}
