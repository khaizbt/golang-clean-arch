package main

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306/go_mux)")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	router := mux.NewRouter()

	// post := router.PathPrefix("post")
	router.HandleFunc("post", getPost).Methods("GET")
	http.ListenAndServe(":8000", router)
}

func getPost(w http.ResponseWriter, r *http.Request) {

}
