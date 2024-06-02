package handlers

import (
	layout "goHtmx/web/templates/base_layout"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {

	return layout.Login().Render(c.Request().Context(), c.Response())
}
