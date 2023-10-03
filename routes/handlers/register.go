package handlers

import (
	"rearrange/app/register/controller"
	"rearrange/app/register/repository"
	"rearrange/app/register/service"
	"rearrange/package/database"

	"github.com/labstack/echo/v4"
)

type handlersRegister struct {
	Controller controller.Controller
}

func NewHandlerAdmin()*handlersRegister {
	mdr := repository.NewRepository()
	mds := service.NewService(database.DBManager(), mdr)

	return &handlersRegister{
		Controller: controller.NewController(mds),
	}
}

func (hd *handlersRegister) Route(g *echo.Group) {
	g.GET("", hd.Controller.GetAll)	
	g.GET("/:id", hd.Controller.GetByID)	
	g.POST("", controller.CreateAdmin)
	g.PUT("/:id", controller.UpdateAdmin)
	g.DELETE("/:id", controller.DeleteAdmin)
}