package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func tasksHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, "Get method")
	case http.MethodPost:
		fmt.Fprint(w, "Post method")
	}
}

func main() {
	InitDB()
	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", tasksHandlers).Methods(http.MethodGet, http.MethodPost)
	http.ListenAndServe(":8080", router)
}
