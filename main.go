package main

import (
	"rearrange/package/database"
	"rearrange/package/log"
	"rearrange/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	database.InitDB()
	database.Migrate()

	routes.SetupRoutes(e.Group(""))
	e.Use(log.LogRequest)
	e.Logger.Fatal(e.Start(":8000"))
}