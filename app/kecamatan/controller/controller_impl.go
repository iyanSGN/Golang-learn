package controller

import (
	"fmt"
	"net/http"
	"rearrange/app/kecamatan"
	"rearrange/app/kecamatan/repository"
	"rearrange/app/kecamatan/service"
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
	ID, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		return response.ErrorResponse(c, response.BuildError(response.ErrBadRequest, err))
	}

	result, err := co.Service.GetByID(c,uint(ID))
	if err != nil {
		return response.ErrorResponse(c, err )
	}

	return response.SuccessResponse(c, http.StatusOK,"Success Get Kecamatan by Id", result)

}

func CreateKecamatan(c echo.Context) error {
	var kec kecamatan.KecamatanRequestDTO
	if err := c.Bind(&kec); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	CreatedKecamatan, err := repository.CreateKecamatan(kec)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status_code" : http.StatusCreated,
		"data" : CreatedKecamatan,
	})
}

func UpdateKecamatan(c echo.Context) error {
	kecID := c.Param("id")
	id, err :=  strconv.Atoi(kecID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updatedKecamatan := kecamatan.KecamatanRequestDTO{}
	if err := c.Bind(&updatedKecamatan);
	err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = repository.UpdateKecamatan(id, updatedKecamatan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "Kecamatan Updated Successfully",
		"status_code" : http.StatusOK,
	})
}

func DeleteKabupaten(c echo.Context) error {
	kecID :=  c.Param("id")
	id, err := strconv.Atoi(kecID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid id"))
	}

	err = repository.DeleteKecamatan(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "Kecamatan has been deleted successfully",
		"status_code" : http.StatusOK,
	})
}