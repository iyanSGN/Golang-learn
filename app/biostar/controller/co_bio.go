package controller

import (
	"net/http"
	"rearrange/app/biostar/repository"

	"github.com/labstack/echo/v4"
)

func HandleUser(c echo.Context) error {
	userID, err := repository.GetUser()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":        userID,
		"status_Code": http.StatusOK,
	})
}


func HandlePost(c echo.Context) error {
	var request repository.Post

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	createdUser, err := repository.PostUser(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"status_code" : http.StatusOK,
		"data" : map[string]interface{}{
			"message" :  "register Successfull",
			"Name" : createdUser,
		},
	}

	return c.JSON(http.StatusOK, response)
}
