package service

import (
	"rearrange/app/register"

	"github.com/labstack/echo/v4"
)

type Service interface {
	GetAll(c echo.Context) ([]register.AdminResponseDTO, error)
	GetByID(c echo.Context, id uint) (register.AdminResponseDTO, error)
}