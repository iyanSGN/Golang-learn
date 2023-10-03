package service

import (
	"rearrange/app/provinsi"

	"github.com/labstack/echo/v4"
)

type Service interface {
	GetAll(c echo.Context) ([]provinsi.ProvinsiResponseDTO, error)
	GetByID(c echo.Context, ID uint) (provinsi.ProvinsiResponseDTO, error)
}