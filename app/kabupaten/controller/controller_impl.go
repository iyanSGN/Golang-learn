package controller

import (
	"fmt"
	"net/http"
	"rearrange/app/kabupaten"
	"rearrange/app/kabupaten/repository"
	"rearrange/app/kabupaten/service"
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

	return response.SuccessResponse(c, http.StatusOK, "Success Get All Kabupaten", result)
}

func (co *controllerImpl)GetByID(c echo.Context) error {
	strID := c.Param("id")

	ID, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		return response.ErrorResponse(c, response.BuildError(response.ErrNotFound, err))
	}

	result, err := co.Service.GetByID(c, uint(ID)) 
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, http.StatusOK, "Success Get Kabupaten by id", result)
}

func CreateKabupaten(c echo.Context) error {
	var kab models.MKabKota
	if err := c.Bind(&kab);
	err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	CreatedKabupaten, err := repository.CreateKabupaten(kab)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"data" : map[string]interface{}{
			"id" : CreatedKabupaten.ID,
			"Provinsi_id" : CreatedKabupaten.IDProvinsi,
			"Kabupaten_nama" : CreatedKabupaten.Nama,
		},
	}
	
	return c.JSON(http.StatusOK, response)
}

func UpdateKabupaten(c echo.Context) error {
	kabID := c.Param("id")
	id, err := strconv.Atoi(kabID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid id"))
	}

	UpdatedKabupaten := kabupaten.KabKotaRequestDTO{}
	if err := c.Bind(&UpdatedKabupaten); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = repository.UpdateKabupaten(id, UpdatedKabupaten)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Kabupaten Updated Successfully",
		"status_code": http.StatusOK,
	})


}

func DeleteKabupaten(c echo.Context) error {
	kabID := c.Param("id")
	id, err := strconv.Atoi(kabID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid user id"))
	}

	err = repository.DeleteKabupaten(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "Kabupaten deleted successfully",
		"status_code" : http.StatusOK,
	})
}

