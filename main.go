package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func connectDatabase() (db *sql.DB) {
	//Open db connection
	db, err := sql.Open("postgres", "user=postgres password=password dbname=noteBookApp sslmode=disable")

	if err != nil {
		log.Panic(err)
	}

	return db
}

func main() {
	db := connectDatabase()
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/login", login).Methods("GET")
}

// LOGIN
func login(w http.ResponseWriter, r *http.Request) {

}


// MODEL USER
type User struct {
	user
}
