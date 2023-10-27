package controller

import "github.com/labstack/echo/v4"

type Controller interface {
	GetAll(c echo.Context) error
	GetByID(c echo.Context) error
}