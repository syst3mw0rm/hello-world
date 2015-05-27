package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

var (
	DB_USER     = os.GetEnv("DB_USER")
	DB_PASSWORD = os.GetEnv("DB_PASSWORD")
	DB_NAME     = os.GetEnv("DB_NAME")
	DB_HOST     = os.GetEnv("DB_HOST")
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/record", record_tx)
	http.ListenAndServe(":8000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
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
	err = db.QueryRow("INSERT INTO tx(from_user, to_user) values ('a', 'b') returning uid;").Scan(&lastInsertId)
	if err != nil {
		panic(err)
	}

	fmt.Println("last inserted id =", lastInsertId)
}
