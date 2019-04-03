package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/than-os/sentinel-bot/nodes/master-node/service"
)


func main() {

	e := echo.New()

	//middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE},
	}))

	e.GET("/", service.RootFunc)
	e.POST("/user", service.AddNewUser)
	e.DELETE("/user", service.RemoveUser)

	//Start the server
	e.Start(":30002")
}