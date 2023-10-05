package controller

import (
	"fmt"
	"net/http"
	"rearrange/app/register/repository"
	"rearrange/app/register/service"
	"rearrange/helpers"
	"rearrange/middleware"
	"rearrange/models"
	"rearrange/package/response"
	"rearrange/token"
	"strconv"

	"github.com/labstack/echo/v4"
)

type controllerImpl struct {
	Service service.Service
}

func NewController(Service service.Service) Controller {
	return &controllerImpl{
		Service: Service,
	}
}

func (co *controllerImpl)GetAll(c echo.Context) error {
	result, err := co.Service.GetAll(c)
	if err != nil {
		return response.ErrorResponse(c,err)
	}

	return response.SuccessResponse(c, http.StatusOK, "Success Get All Admin", result)
}

func (co *controllerImpl)GetByID(c echo.Context) error {
	adminID := c.Param("id")

	id, err := strconv.ParseUint(adminID, 10, 64)
	if err != nil {
		return response.ErrorResponse(c,response.BuildError(response.ErrBadRequest, err))
	}

	result, err :=  co.Service.GetByID(c, uint(id))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, http.StatusOK, "Success Get Admin by ID", result)


}

func CreateAdmin(c echo.Context) error {
	var admin models.MRegister
	otpRepo := middleware.NewOtpRepository()

	if err := c.Bind(&admin);
	err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":     "Bad Request",
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
		})
	}

	admin.Isactive = 0

	if err := helpers.HashPassword(&admin);
	err != nil {
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
		"status_code" : http.StatusOK,
		"token" : token,
		"data" : map[string]interface{}{
			"message" :  "register Successfull",
			"email" : createdAdmin.Email,
			"password" : createdAdmin.Password,
			"isactive" : createdAdmin.Isactive,
		},
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateAdmin(c echo.Context) error {
	adminID := c.Param("id")
	id, err := strconv.Atoi(adminID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updatedAdmin := models.MRegister{}
	if err := c.Bind(&updatedAdmin);
	err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if updatedAdmin.Password != "" {
		if err := helpers.HashPassword(&updatedAdmin);
		err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

	}

	err = repository.UpdateAdmin(id,updatedAdmin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "Admin has successfully updated",
		"status_code" :	http.StatusOK,
	})
}


func DeleteAdmin(c echo.Context) error {
	adminID := c.Param("id")
	id, err := strconv.Atoi(adminID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid user id"))
	}

	err = repository.DeleteAdmin(id)
	if err != nil {
	return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "Selected user has been deleted",
		"status_code" : http.StatusOK,
	})
}