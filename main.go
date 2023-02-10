package main

import (
	"teguh/factory"
	"teguh/utils/database/mongoDB"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	client := mongoDB.InitDB()
	e := echo.New()
	factory.InitFactory(e, client)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(":8080"))

	// Enable logging
	e.Logger.SetLevel(log.DEBUG)
}
