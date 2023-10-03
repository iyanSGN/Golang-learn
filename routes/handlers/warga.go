package handlers

import (
	"rearrange/app/warga/controller"
	"rearrange/app/warga/repository"
	"rearrange/app/warga/service"
	"rearrange/package/database"

	"github.com/labstack/echo/v4"
)


type handlersWarga struct {
	Controller controller.Controller
}

func NewHandlerWarga() *handlersWarga {
	mdr := repository.NewRepo()
	mds := service.NewService(database.DBManager(), mdr)

	return &handlersWarga{
		Controller: controller.NewController(mds),
	}
}

func(hd *handlersWarga) Route(g *echo.Group) {
	g.GET("", hd.Controller.GetAll)
	g.GET("/:id", hd.Controller.GetByID)
	g.POST("", controller.CreateWarga)
	g.PUT("/:id", controller.UpdateWarga)
	g.DELETE("/:id", controller.DeleteWarga)
}