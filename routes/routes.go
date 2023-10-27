package routes

import (
	"rearrange/app/register/control"
	"rearrange/app/biostar/controller"

	"rearrange/routes/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(g *echo.Group) {
	g.GET("/", handlers.Home)
	g.POST("/login", handlers.LoginAccount)
	g.POST("/generateotp", handlers.GenerateOTP)
	g.POST("/resendotp", handlers.ResendOTP)
	g.POST("/verifyotp", handlers.VerifyOtp)
	g.POST("/admin", control.CreateAdmin)

	
	Admin := g.Group("")
	Admin.Use(handlers.TokenMiddleware)

	//BIOSTAR API
	Admin.GET("/biostar", controller.HandleUser)
	Admin.POST("/biostar", controller.HandlePost)

	kecamatanGroup := Admin.Group("/kecamatan")
	kabupatenGroup := Admin.Group("/kabupaten")
	provinsiGroup := Admin.Group("/provinsi")
    wargaGroup := Admin.Group("/warga")

	
	
	// Routes
	handlers.NewHandlerUser().Route(g.Group("/user"))
	handlers.NewHandlerKecamatan().Route(kecamatanGroup)
	handlers.NewHandlerKabupaten().Route(kabupatenGroup)
    handlers.NewHandlerProvinsi().Route(provinsiGroup)
    handlers.NewHandlerWarga().Route(wargaGroup)
}

