package config

import (
	"assignment-2/helpers"
	"assignment-2/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() {
	helpers.LoadEnvChecking()

	var (
		host     = os.Getenv("PG_HOST")
		user     = os.Getenv("PG_USER")
		password = os.Getenv("PG_PASSWORD")
		dbName   = os.Getenv("PG_DBNAME")
		dbPort   = os.Getenv("PG_PORT")
		db       *gorm.DB
		err      error
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	db.AutoMigrate(models.Order{}, models.Item{})
}
