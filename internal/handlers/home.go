package handlers

import (
	"goHtmx/web/templates/home"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	return render(c, home.HomePage())
}
