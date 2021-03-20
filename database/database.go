package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
	// "github.com/jerrydevin96/lifo-queue/config"
)

const (
	host     = "35.238.100.247"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "project"
)

func connectDB() (*sql.DB, error) {
	var db *sql.DB
	var err error
	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	config.Configurations.DBHost, config.Configurations.DBPort,
	// 	config.Configurations.DBUser, config.Configurations.DBPassword,
	// 	config.Configurations.DBName)
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("Connected to DB")

	return db, err
}

func GetLastRecord() (int, string, error) {
	var lastIndex int
	var lastValue string
	var err error
	log.Println("fetching last record from lifo table")
	query := `select "entry", "value" from lifo order by entry desc limit 1`
	db, err := connectDB()
	if err != nil {
		log.Println(`[ERROR occured] ` + err.Error())
		return 0, "", err
	}
	defer db.Close()
	log.Println("executing query")
	rows, err := db.Query(query)
	if err != nil {
		log.Println(`[ERROR occured] ` + err.Error())
		return 0, "", err
	}

	defer rows.Close()
	rows.Next()
	err = rows.Scan(&lastIndex, &lastValue)
	if err != nil {
		log.Println(`[ERROR occured] ` + err.Error())
		return 0, "", err
	}
	log.Println("Query response for last record entry = " + strconv.Itoa(lastIndex) + " value = " + lastValue)
	return lastIndex, lastValue, err
}

func DeleteLastRecord(entry int) error {
	var err error
	deleteStatement := `delete from lifo where entry=` + strconv.Itoa(entry)
	db, err := connectDB()
	if err != nil {
		log.Println(`[ERROR occured] ` + err.Error())
		return err
	}
	defer db.Close()
	log.Println("deleting last record from table")
	_, err = db.Exec(deleteStatement)
	if err != nil {
		log.Println(`[ERROR occured] ` + err.Error())
		return err
	}
	log.Println("delete successful")
	return err
}

func InsertNewRecord(entry int, value string) error {
	var err error
	db, err := connectDB()
	if err != nil {
		log.Println(`[ERROR occured] ` + err.Error())
		return err
	}
	defer db.Close()
	log.Println("inserting new record into lifo table")
	insertStatement := `insert into lifo ("entry", "value") values ($1, $2)`
	_, err = db.Exec(insertStatement, entry, value)
	if err != nil {
		log.Println(`[ERROR occured] ` + err.Error())
		return err
	}
	log.Println("insert successful")
	return err
}