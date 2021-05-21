package main

import (
	"net/http"

	"github.com/Makinami/oapi-codegen-bug/app"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	g := e.Group("/api/v1")

	service := &app.AppService{}

	app.RegisterHandlers(g, service)

	if err := e.Start(":4000"); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
