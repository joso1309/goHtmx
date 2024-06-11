package handlers

import (
	"goHtmx/internal/models"
	"goHtmx/web/templates/users"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetUsers(c echo.Context) error {
	var items []models.User
	for i := 0; i < 5; i++ {
		items = append(items, fetchUser(1))
	}
	return render(c, users.UsersTable(items))
}

func GetUser(c echo.Context) error {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	user := fetchUser(id)
	form := users.UserForm(user)

	return render(c, form)
}

func EditUser(c echo.Context) error {

	var user models.User
	c.Bind(&user)

	log.Info("name:", user.FirstName+" "+user.LastName, " select:", user.Select, " datetime:", user.Date+" "+user.Time)
	return GetUsers(c)
}

func fetchUser(id int) models.User {
	return models.User{
		ID:        id,
		FirstName: "Josip",
		LastName:  "Razov",
		Email:     "joso.razov@gmail.com",
	}
}
