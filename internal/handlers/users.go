package handlers

import (
	database "goHtmx/internal/db"
	"goHtmx/internal/types"
	"goHtmx/web/templates/users"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetUsers(c echo.Context) error {
	ctx := c.Get("dbCtx").(*database.DbContext)

	models, err := ctx.GetUsers()
	if err != nil {
		return err
	}

	var items []types.User

	for _, v := range *models {
		item := types.User{
			ID:        v.ID,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Email:     v.Email,
		}
		items = append(items, item)
	}

	return render(c, users.UsersTable(items))
}

func GetUser(c echo.Context) error {
	ctx := c.Get("dbCtx").(*database.DbContext)

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return err
	}

	model, err := ctx.GetUserByID(uint(id))
	if err != nil {
		return err
	}

	user := types.User{
		ID:        model.ID,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Email:     model.Email,
	}
	form := users.UserForm(user)

	return render(c, form)
}

func EditUser(c echo.Context) error {
	var user types.User
	c.Bind(&user)

	log.Info("name:", user.FirstName+" "+user.LastName, " select:", user.Select, " datetime:", user.Date+" "+user.Time)
	return GetUsers(c)
}
