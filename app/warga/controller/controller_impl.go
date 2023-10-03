package controller

import (
	"fmt"
	"net/http"
	"rearrange/app/warga"
	"rearrange/app/warga/repository"
	"rearrange/app/warga/service"
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

	return response.SuccessResponse(c, http.StatusOK, "Success Get All Warga", result)
}

func (co *controllerImpl) GetByID(c echo.Context) error {
	strID := c.Param("id")

	ID, err := strconv.ParseUint(strID,10,64)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	result, err := co.Service.GetByID(c, uint(ID)) 
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, http.StatusOK,"Success Get Warga by Id", result)


}

func CreateWarga(c echo.Context) error {
	var warga models.MWarga
	if err := c.Bind(&warga);
	err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	createdWarga, err := repository.CreateWarga(warga)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"data" : map[string]interface{}{
			"id" : createdWarga.ID,
			"nama" : createdWarga.Nama,
			"no_ktp" : createdWarga.NoKtp,
			"kecamatan_id" : createdWarga.IDKecamatan,
			"kabupaten_id" : createdWarga.IDKabupaten,
			"provinsi_id"  : createdWarga.IDProvinsi,
		},
	}

	return c.JSON(http.StatusOK, response)
}


func UpdateWarga(c echo.Context) error {
	wargaID := c.Param("id")
	id, err :=  strconv.Atoi(wargaID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid id"))
	}

	updatedWarga := warga.WargaRequestDTO{}
	if err :=  c.Bind(&updatedWarga);
	err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = repository.UpdateWarga(id, updatedWarga)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "Warga has been updated",
		"status_code" : http.StatusOK,
	})
}


func DeleteWarga(c echo.Context) error {
	wargaID := c.Param("id")
	id, err :=  strconv.Atoi(wargaID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid user id"))
	}

	err = repository.DeleteWarga(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "Warga has been deleted",
		"status_code" : http.StatusOK,
	})
} 