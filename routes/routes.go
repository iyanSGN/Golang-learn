package routes

import (
	"rearrange/routes/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(g *echo.Group) {
	g.GET("/", handlers.Home)
	g.POST("/login", handlers.LoginAccount)
	Admin := g.Group("")
	Admin.Use(handlers.TokenMiddleware)

	kecamatanGroup := Admin.Group("/kecamatan")
	kabupatenGroup := Admin.Group("/kabupaten")
	provinsiGroup := Admin.Group("/provinsi")
    wargaGroup := Admin.Group("/warga")
	
	
	// Routes
	handlers.NewHandlerAdmin().Route(g.Group("/admin"))
	handlers.NewHandlerKecamatan().Route(kecamatanGroup)
	handlers.NewHandlerKabupaten().Route(kabupatenGroup)
    handlers.NewHandlerProvinsi().Route(provinsiGroup)
    handlers.NewHandlerWarga().Route(wargaGroup)
}

