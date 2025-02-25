package db

import (
	"log"

	"github.com/mrzhov/course-app/pkg/task"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn string = "host=localhost user=cource_user password=cource_password dbname=cource_app port=5432 sslmode=disable"

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to database: ", err)
	}

	db.AutoMigrate(&task.Task{})

	return db
}
