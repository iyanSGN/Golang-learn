package service

import (
	"rearrange/app/kabupaten"

	"github.com/labstack/echo/v4"
)

type Service interface {
	GetAll(c echo.Context) ([]kabupaten.KabKotaResponseDTO, error)
	GetByID(c echo.Context, ID uint) (kabupaten.KabKotaResponseDTO, error)
}