package main

import (
	"goHtmx/internal/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano} method=${method}, uri=${uri}, status=${status}, error:${error}\n",
	}))
	e.Use(middleware.Recover())

	// e.Static("/static", "../../web/static") // on debug
	// e.Static("/static", "static") // on prod
	e.Static("/static", "web/static") // on air

	RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", handlers.Home)

	e.GET("/users", handlers.GetUsers)
	e.GET("/users/:id", handlers.GetUser)
	e.POST("/users", handlers.EditUser)

	e.GET("/login", handlers.Login)
}
