package webserver

import (
	"net/http"
	"strconv"

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

	app.POST("/api/reset", func(c echo.Context) error {
		e.Cpu.Reset()
		return c.Render(http.StatusOK, "content.html", e)
	})

	app.POST("/api/continue-unimplemented", func(c echo.Context) error {
		e.Cpu.ContinueUnimpl()
		return c.Render(http.StatusOK, "content.html", e)
	})

	app.POST("/api/continue-until", func(c echo.Context) error {
		bp, err := strconv.ParseUint(c.FormValue("bp"), 16, 16)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid breakpoint")
		}

		e.Cpu.ContinueUntil(uint16(bp))

		return c.Render(http.StatusOK, "content.html", e)
	})

	app.Static("/static", "static")

	return app
}
