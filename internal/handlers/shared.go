package handlers

import (
	layout "goHtmx/web/templates/base_layout"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {

	if isHXRequest(c) {
		return component.Render(c.Request().Context(), c.Response())
	}
	app := layout.App(component)
	return app.Render(c.Request().Context(), c.Response())
}

func isHXRequest(c echo.Context) bool {
	return c.Request().Header.Get("HX-Request") == "true"
}
