package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

var (
	DB_USER     = os.Getenv("DB_ENV_DB_USER")
	DB_PASSWORD = os.Getenv("DB_ENV_DB_PASSWORD")
	DB_NAME     = os.Getenv("DB_ENV_DB_NAME")
	DB_HOST     = os.Getenv("DB_PORT_5432_TCP_ADDR")
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/debug", debug)
	http.HandleFunc("/record", record_tx)
	http.HandleFunc("/all_tx", record_tx)
	http.ListenAndServe(":8000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func debug(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, os.Environ())
}

func record_tx(w http.ResponseWriter, r *http.Request) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME, DB_HOST)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var lastInsertId int

	fromUser := req.FormValue("from_user")
	toUser := req.FormValue("to_user")

	err = db.QueryRow("INSERT INTO tx(from_user, to_user) values ($1, $2) returning id;", fromUser, toUser).Scan(&lastInsertId)
	if err != nil {
		panic(err)
	}

	fmt.Println("last inserted id =", lastInsertId)
}

func all_tx(w http.ResponseWriter, r *http.Request) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME, DB_HOST)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tx;")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "%#v\n", rows)
}
