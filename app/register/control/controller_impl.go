package control

import (
	"net/http"
	"rearrange/app/register/repository"
	"rearrange/helpers"
	"rearrange/middleware"
	"rearrange/models"
	"rearrange/token"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateAdmin(c echo.Context) error {
	var admin models.MRegister
	otpRepo := middleware.NewOtpRepository()

	if err := c.Bind(&admin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":     "Bad Request",
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
		})
	}

	admin.Isactive = 0

	if err := helpers.HashPassword(&admin); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}

	createdAdmin, err := repository.CreateAdmin(admin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":     "Failed to create user",
			"error":       err.Error(),
			"status_code": http.StatusInternalServerError,
		})
	}

	otpCode := otpRepo.GenerateOTP()
	err = otpRepo.CreateOtp(int32(createdAdmin.ID), otpCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	otpcodeStr := strconv.Itoa(int(otpCode))

	err = middleware.Otp(createdAdmin.Email, otpcodeStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	adminID := uint(admin.ID)
	token, err := token.GenerateToken(adminID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"status_code": http.StatusOK,
		"token":       token,
		"data": map[string]interface{}{
			"message":  "register Successfull",
			"email": createdAdmin.Email,
			"password": createdAdmin.Password,
			"isactive": createdAdmin.Isactive,
		},
	}

	return c.JSON(http.StatusOK, response)
}
