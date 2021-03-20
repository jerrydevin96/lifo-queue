package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "35.238.100.247"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "project"
)

func main() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
	response, err := db.Exec("create table lifo(entry int unique, value varchar(255))")
	CheckError(err)
	fmt.Println(response)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
