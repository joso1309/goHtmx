package handlers

import (
	layout "goHtmx/web/templates/base_layout"
	"goHtmx/web/templates/home"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type LoginData struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func Login(c echo.Context) error {

	return layout.Login(false).Render(c.Request().Context(), c.Response())
}

func Logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1 // This will cause the session cookie to expire
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusFound, "/login")
}

func Authorize(c echo.Context) error {

	var login LoginData
	err := c.Bind(&login)
	if err != nil {
		return err
	}
	log.Info("user:", login.Username, " pwd:", login.Password)

	if login.Username != "da" || login.Password != "da" {

		return layout.Login(true).Render(c.Request().Context(), c.Response())
	}

	sess, _ := session.Get("session", c)
	sess.Values["authenticated"] = true
	sess.Save(c.Request(), c.Response())

	app := layout.App(home.HomePage())
	c.Response().Header().Set("HX-Replace-Url", "/")
	return app.Render(c.Request().Context(), c.Response())
}
