package handlers

import (
	"rearrange/app/kabupaten/controller"
	"rearrange/app/kabupaten/repository"
	"rearrange/app/kabupaten/service"
	"rearrange/package/database"

	"github.com/labstack/echo/v4"
)

type handlersKabupaten struct {
	Controller controller.Controller
}

func NewHandlerKabupaten() *handlersKabupaten {
	mdr := repository.NewRepository()
	mds := service.NewService(database.DBManager(), mdr)

	return &handlersKabupaten{
		Controller: controller.NewController(mds),
	}
}

func (hd *handlersKabupaten) Route(g *echo.Group) {
	g.GET("", hd.Controller.GetAll )
	g.GET("/:id", hd.Controller.GetByID)
	g.POST("", controller.CreateKabupaten)
	g.PUT("/:id", controller.UpdateKabupaten)
	g.DELETE("/:id", controller.DeleteKabupaten)
}