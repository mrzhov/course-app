package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn string = "host=localhost user=cource_user password=cource_password dbname=cource_app port=5432 sslmode=disable"

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to database: ", err)
	}

	return db
}
