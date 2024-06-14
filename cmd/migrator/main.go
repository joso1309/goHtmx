package main

import (
	database "goHtmx/internal/db"
	"goHtmx/internal/shared"
	"log"
)

func main() {
	err := shared.LoadEnv()
	if err != nil {
		log.Fatal("Error loading .env file", "err", err)
	}

	ctx, err := database.NewDbContext()
	if err != nil {
		log.Fatal(err)
	}

	database.InitilizeDB(*ctx)
	// database.SeedDB(*ctx)
}
