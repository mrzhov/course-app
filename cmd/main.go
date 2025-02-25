package main

import (
	"net/http"

	"github.com/gorilla/mux"
	DB "github.com/mrzhov/course-app/pkg/common/db"
	"github.com/mrzhov/course-app/pkg/task"
)

func main() {
	r := mux.NewRouter()
	db := DB.Init()

	task.Controller(r, db)

	http.ListenAndServe(":8080", r)
}
