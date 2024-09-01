package main

import (
	"github.com/duch94/meh_cms/src/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	router.InitRouter(e)

	err := e.Start(":8080")
	e.Logger.Fatal(err)
}
