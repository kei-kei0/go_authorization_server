package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Debug = true
	e.Logger.SetLevel(log.DEBUG)
	e.GET("/ping", func(c echo.Context) error {
		e.Logger.Debug("pong!")
		return c.String(http.StatusOK, "pong!")
	})
	e.GET("/", func(c echo.Context) error {
		e.Logger.Debug("Hello World!!")
		return c.String(http.StatusOK, "Hello World!")
	})
	e.Logger.Fatal(e.Start(":8888"))
}
