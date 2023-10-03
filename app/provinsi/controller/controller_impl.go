package controller

import (
	"fmt"
	"net/http"
	"rearrange/app/provinsi/repository"
	"rearrange/app/provinsi/service"
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

func (co *controllerImpl) GetAll(c echo.Context) error {
	result, err := co.Service.GetAll(c)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, http.StatusOK, "Success Get All Kecamatan", result)
}

func (co *controllerImpl) GetByID(c echo.Context) error {
	strID := c.Param("id")
	ID, err := strconv.ParseUint(strID,10,64)
	if err != nil {
		return response.ErrorResponse(c, response.BuildError(response.ErrBadRequest,err))
	}

	result, err := co.Service.GetByID(c, uint(ID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, http.StatusOK, "Success get id Provinsi", result)
}

func CreateProvinsi(c echo.Context) error {
	request := models.MProvinsi{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	CreatedProvinsi, err := repository.CreateProvinsi(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"data" : map[string]interface{}{
			"id" : CreatedProvinsi.ID,
			"nama" : CreatedProvinsi.Nama,
		},
		}
		return c.JSON(http.StatusOK, response)
	}


func UpdateProvinsi(c echo.Context) error {
	provID := c.Param("id")
	id, err := strconv.Atoi(provID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid id"))
	}

	updatedProvinsi := models.MProvinsi{}
	if err := c.Bind(&updatedProvinsi);
	err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = repository.UpdateProvinsi(id, updatedProvinsi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "User Updated Successfully",
		"status_code": http.StatusOK,
	})
}

func DeleteUser(c echo.Context) error {
	provID := c.Param("id")

	id, err := strconv.Atoi(provID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid id"))
	}

	err = repository.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "User id  has been deleted",
		"status_code" : http.StatusOK,
	})
}