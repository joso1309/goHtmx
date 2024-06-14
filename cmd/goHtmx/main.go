package main

import (
	database "goHtmx/internal/db"
	"goHtmx/internal/handlers"
	"goHtmx/internal/shared"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// e.Static("/static", "../../web/static") // on debug
	// e.Static("/static", "static") // on prod
	e.Static("/static", "web/static") // on air

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano} method=${method}, uri=${uri}, status=${status}, error:${error}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	err := shared.LoadEnv()
	if err != nil {
		e.Logger.Fatal("Error loading .env file", "err", err)
	}

	dbCtx, err := database.NewDbContext()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("dbCtx", dbCtx)
			return next(c)
		}
	})

	RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", handlers.Home, isAuthenticated)

	e.GET("/users", handlers.GetUsers, isAuthenticated)
	e.GET("/users/:id", handlers.GetUser, isAuthenticated)
	e.POST("/users", handlers.EditUser, isAuthenticated)

	e.GET("/login", handlers.Login)
	e.POST("/login", handlers.Authorize)
	e.GET("/logout", handlers.Logout)
}

func isAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		if auth, ok := sess.Values["authenticated"].(bool); !ok || !auth {
			return c.Redirect(http.StatusFound, "/login")
		}
		return next(c)
	}
}
