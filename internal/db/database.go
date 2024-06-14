package database

import (
	"errors"
	"log"
	"log/slog"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DbContext struct {
	Db *gorm.DB
}

func NewDbContext() (*DbContext, error) {
	connString := os.Getenv("CONN_STRING")
	if connString == "" {
		return nil, errors.New("db conn string missing")
	}

	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DbContext{Db: db}, nil
}

func InitilizeDB(ctx DbContext) {

	err := ctx.Db.AutoMigrate(&User{})
	if err != nil {
		slog.Error("Error on migration", "err", err)
	}
}

func SeedDB(ctx DbContext) {

	user := User{
		FirstName: "Joso",
		LastName:  "Ražov",
		Email:     "joso.razov@gmail.com",
	}

	result := ctx.Db.Create(&user)

	if result.Error != nil {
		log.Fatal(result.Error.Error())
	}
}
