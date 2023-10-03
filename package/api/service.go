package api

import (
	"net/http"
	"rearrange/package/response"

	"github.com/labstack/echo/v4"
)

func GetToken(c echo.Context) (string, error) {
	authToken := c.Request().Header.Get("Authorization")
	if authToken == "" {
		return "", response.BuildCustomError(http.StatusUnauthorized, "Unauthorized")
	}

	return authToken, nil
}