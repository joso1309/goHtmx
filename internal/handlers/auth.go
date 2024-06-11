package handlers

import (
	layout "goHtmx/web/templates/base_layout"
	"goHtmx/web/templates/home"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type LoginData struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func Login(c echo.Context) error {

	return layout.Login().Render(c.Request().Context(), c.Response())
}

func Authorize(c echo.Context) error {

	var login LoginData
	err := c.Bind(&login)
	if err != nil {
		return err
	}
	log.Info("user:", login.Username, " pwd:", login.Password)

	app := layout.App(home.HomePage())
	return app.Render(c.Request().Context(), c.Response())
}
