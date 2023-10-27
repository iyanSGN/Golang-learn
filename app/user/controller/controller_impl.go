package controller


import (
	"fmt"
	"net/http"
	"rearrange/app/user/repository"
	"rearrange/app/user/service"
	"rearrange/helpers"
	"rearrange/models"
	"rearrange/package/response"
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