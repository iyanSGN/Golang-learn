package controller

import (
	// "fmt"
	// "strconv"

	"fmt"
	"net/http"
	"rearrange/app/biostar/repository"

	// "rearrange/app/biostar/repository"

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
	var request repository.Users

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


func DeleteUser(c echo.Context) error {

	err := repository.DeleteUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "User deleted successfully",
		"status_code" : http.StatusOK,
	})
}

func HandleLogin(c echo.Context) error {
	var request repository.Login

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	loginResponse, headers, err := repository.LoginAdmin(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Set the bs-session-id cookie in the response
	sessionID := headers.Get("bs-session-id")
	if sessionID != "" {
		cookie := new(http.Cookie)
		cookie.Name = "bs-session-id"
		cookie.Value = sessionID
		c.SetCookie(cookie)
	}

	fmt.Println("================================================")
	fmt.Println(sessionID)
	fmt.Println("================================================")

	return c.JSON(http.StatusOK, loginResponse)
}


func HandleLogout(c echo.Context) error {
	logout, err := repository.LogoutAdmin()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":        logout,
		"status_Code": http.StatusOK,
	})
}
