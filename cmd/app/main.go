package main

import (
	"net/http"

	"github.com/gorilla/mux"
	DB "github.com/mrzhov/course-app/internal/common/db"
	"github.com/mrzhov/course-app/internal/task"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := mux.NewRouter()
	db := DB.Init(dbUrl)

	task.RegisterRoutes(r, db)

	http.ListenAndServe(port, r)
}
