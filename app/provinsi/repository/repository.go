package repository

import (
	"rearrange/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(c echo.Context, DB *gorm.DB) ([]models.MProvinsi, error)
	GetByID(c echo.Context, DB *gorm.DB, ID uint) (models.MProvinsi, error)
}