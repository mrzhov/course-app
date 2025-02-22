package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var task string

type requestBody struct {
	Task string `json:"task"`
}

func handleError(w io.Writer, message string) {
	fmt.Fprintf(w, "Error: %s", message)
}

func helloHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if len(task) == 0 {
			handleError(w, "No task. Use POST route first.")
			return
		}
		fmt.Fprintf(w, "Hello, %s!", task)
	case "POST":
		body := requestBody{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			handleError(w, err.Error())
			return
		}
		if len(body.Task) == 0 {
			handleError(w, "Field \"task\" is required in body!")
			return
		}
		task = body.Task
		fmt.Fprint(w, "Task added successfully! Check GET route.")
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/hello", helloHandlers).Methods("GET", "POST")

	http.ListenAndServe(":8080", router)
}
