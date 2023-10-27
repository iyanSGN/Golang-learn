package handlers

import (
	"rearrange/app/user/controller"
	"rearrange/app/user/repository"
	"rearrange/app/user/service"
	"rearrange/package/database"

	"github.com/labstack/echo/v4"
)

type handlersUser struct {
	Controller controller.Controller
}

func NewHandlerUser() *handlersUser {
	mdr := repository.NewRepo()
	mds := service.NewService(database.DBManager(), mdr)

	return &handlersUser{
		Controller: controller.NewController(mds),
	}
}

func (hd *handlersUser) Route(g *echo.Group) {
	g.GET("", hd.Controller.GetAll)
	g.GET("/:id", hd.Controller.GetByID)
	g.PUT("/:id", controller.UpdateAdmin)
	g.DELETE("/:id", controller.DeleteAdmin)
}