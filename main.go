package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func connectDatabase() (db *sql.DB) {
	//Open db connection
	db, err := sql.Open("postgres", "user=postgres password=password dbname=noteDB sslmode=disable")

	if err != nil {
		log.Panic(err)
	}

	return db
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/login", login).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

// LOGIN
func login(w http.ResponseWriter, r *http.Request) {

	db := connectDatabase()
	defer db.Close()

	var userLogin User
	req, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(req, &userLogin)

	if checkUsername(userLogin.UserName) {

		usernameCookie := &http.Cookie{ // create a username cookie
			Name:  "username",         // cookie name
			Value: userLogin.UserName, // stored username
		}

		http.SetCookie(w, usernameCookie)  // set user name cookie
		fmt.Fprint(w, "Login Successfull") // print for correct login details

	} else {
		fmt.Fprint(w, "Login Unsuccessfull, bad username") // print for incorrect login details
	}

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

// MODEL USER
type User struct {
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
	UserName   string `json:"userName"`
	Password   string `json:"pasword"`
}
