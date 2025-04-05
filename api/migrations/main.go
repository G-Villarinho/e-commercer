package main

import (
	"context"
	"log"
	"time"

	"github.com/g-villarinho/xp-life-api/config"
	"github.com/g-villarinho/xp-life-api/databases"
	"github.com/g-villarinho/xp-life-api/models"
)

func main() {
	config.LoadEnv()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := databases.NewPostgresDatabase(ctx)
	if err != nil {
		log.Fatal("error to connect to database: ", err)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Session{},
		&models.OTP{},
		&models.Store{},
		&models.Billboard{},
		&models.Category{},
		&models.Size{},
		&models.Color{},
		&models.Product{},
		&models.ProductImage{},
	); err != nil {
		log.Fatal("error to auto migrate: ", err)
	}

	log.Println("migrations done")
}
