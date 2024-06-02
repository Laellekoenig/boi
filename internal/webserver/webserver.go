package webserver

import (
	"net/http"

	"github.com/Laellekoenig/boi/internal/emulator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer(e *emulator.Emulator) *echo.Echo {
	app := echo.New()
	app.Renderer = newTemplate()
	app.Use(middleware.Logger())

	app.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", e)
	})

	app.POST("/api/step", func(c echo.Context) error {
		e.Cpu.ExecuteStep()
		return c.Render(http.StatusOK, "content.html", e)
	})

	app.Static("/static", "static")

	return app
}
