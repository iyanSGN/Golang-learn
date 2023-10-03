package service

import (
	"rearrange/app/kecamatan"

	"github.com/labstack/echo/v4"
)

type Service interface {
	GetAll(c echo.Context) ([]kecamatan.KecamatanResponseDTO, error)
	GetByID(c echo.Context, ID uint) (kecamatan.KecamatanResponseDTO, error)
}