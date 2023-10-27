package service

import (
	"rearrange/app/user"

	"github.com/labstack/echo/v4"
)

type Service interface {
	GetAll(c echo.Context) ([]user.UserResponseDTO, error)
	GetByID(c echo.Context, id uint) (user.UserResponseDTO, error)
}