package service

import (
	service "rearrange/app/warga"

	"github.com/labstack/echo/v4"
)

type Service interface {
	GetAll(c echo.Context) ([]service.WargaResponseDTO, error)
	GetByID(c echo.Context, ID uint) (service.WargaResponseDTO, error)

}