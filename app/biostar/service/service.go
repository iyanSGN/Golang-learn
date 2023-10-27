package service

import (
	"rearrange/app/biostar"

	"github.com/labstack/echo/v4"
)

type Service interface {
	GetAll(c echo.Context) ([]biostar.BioStarResponseDTO, error)
	GetByID(c echo.Context, id uint) (biostar.BioStarResponseDTO, error)
}