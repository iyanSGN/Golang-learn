package handlers

import (
	"fmt"
	"net/http"
	"rearrange/app/login"
	"rearrange/token"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func LoginAccount(c echo.Context) error {
	loginData := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.Bind(&loginData); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid request"))
	}

	admin, err := login.GetAdminByEmail(loginData.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("invalid Email/username, try again"))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(loginData.Password));
	err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	adminID := uint(admin.ID)
	token, err :=  token.GenerateToken(adminID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("failed to generate token"))
	}

	if c.Get("newToken") != nil {
		newToken, ok := c.Get("newToken").(string)
		if ok && newToken != token {
			token = newToken
		}
	}

	response := map[string]interface{}{
		"status_code" : http.StatusOK,
		"token" : token,
		"data" : map[string]interface{}{
			"message" : "Login as an admin successfull",
			"id" : admin.ID,
			"email" : admin.Email,
		},
	}

	return c.JSON(http.StatusOK, response)
}