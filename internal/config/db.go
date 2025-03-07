package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to database: ", err)
	}

	return db
}
