package handlers

import (
	"rearrange/app/provinsi/repository"
	"rearrange/app/provinsi/service"	
	"rearrange/app/provinsi/controller"
	"rearrange/package/database"
	"github.com/labstack/echo/v4"
)

type handlersProvinsi struct {
	Controller controller.Controller
}

func NewHandlerProvinsi()*handlersProvinsi {
	mdr := repository.NewRepository()
	mds	:= service.NewService(database.DBManager(), mdr)

	return &handlersProvinsi{
		Controller: controller.NewController(mds),
	}
}

func (hd *handlersProvinsi)	Route(g *echo.Group) {
	g.GET("", hd.Controller.GetAll)
	g.GET("/:id", hd.Controller.GetByID)
	g.POST("", controller.CreateProvinsi)
	g.PUT("/:id", controller.UpdateProvinsi)
	g.DELETE("/:id", controller.DeleteUser)
}