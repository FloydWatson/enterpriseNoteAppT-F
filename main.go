package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
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

	var userLogin User
	req, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(req, &userLogin)

}

// LOGIN FUNC CHECK USERNAME
func checkUsername(username string) bool {
	var name string

	db := connectDatabase() // db connection
	defer db.Close()        // close db connection after use

	getUserName, err := db.Prepare("Select user_name FROM _user WHERE user_name = $1;") // sql query sent to db $1 is the user name
	if err != nil {
		log.Fatal(err)
	}
	err = getUserName.QueryRow(username).Scan(&name) // sending query

	if err == sql.ErrNoRows {
		return false // return false if username does not exist
	}
	if err != nil {
		log.Fatal(err.Error())
	}

	return true // return true if user name exists
}

// LOGIN FUNC CHECK PASSWORD
func validatePass(password string) bool {
	var pass string

	db := connectDatabase() // db connection
	defer db.Close()        // close db connection after use

	stmt, err := db.Prepare("SELECT password FROM _user WHERE password = $1;") // sql query sent to db $1 is the password
	if err != nil {
		log.Fatal(err)
	}

	err = stmt.QueryRow(password).Scan(&pass) // sending query to db

	if err == sql.ErrNoRows {
		return false // return false if not found
	}
	if err != nil {
		log.Fatal(err)
	}

	return true // return true if correct

}

// MODEL USER
type User struct {
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
	UserName   string `json:"userName"`
	Password   string `json:"pasword"`
}
