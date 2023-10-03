package handlers

import (
	"rearrange/app/kecamatan/controller"
	"rearrange/app/kecamatan/repository"
	"rearrange/app/kecamatan/service"
	"rearrange/package/database"

	"github.com/labstack/echo/v4"
)

type handlersKecamatan struct {
	Controller controller.Controller
}

func NewHandlerKecamatan() *handlersKecamatan {
	mdr := repository.NewRepository()
	mds := service.NewService(database.DBManager(), mdr)

	return &handlersKecamatan{
		Controller: controller.NewController(mds),
	}
}

func (hd *handlersKecamatan) Route(g *echo.Group) {
	g.GET("", hd.Controller.GetAll)
	g.GET("/:id", hd.Controller.GetByID)
	g.POST("", controller.CreateKecamatan)
	g.PUT("/:id", controller.UpdateKecamatan)
	g.DELETE("/:id", controller.DeleteKabupaten)
}