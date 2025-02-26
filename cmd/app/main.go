package main

import (
	"net/http"

	"github.com/gorilla/mux"
	DB "github.com/mrzhov/course-app/internal/common/db"
	"github.com/mrzhov/course-app/internal/task"
)

func main() {
	r := mux.NewRouter()
	db := DB.Init()
	db.AutoMigrate(&task.Task{})

	task.RegisterRoutes(r, db)

	http.ListenAndServe(":8080", r)
}
