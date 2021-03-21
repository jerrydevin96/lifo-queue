package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jerrydevin96/lifo-queue/config"
	_ "github.com/lib/pq"
)

func connectDB() (*sql.DB, error) {
	var db *sql.DB
	var err error
	sqlconn := ``
	sqlDriver := ``
	if config.Configurations.DBProvider == "postgres" {
		sqlDriver = "postgres"
		sqlconn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.Configurations.DBHost, config.Configurations.DBPort,
			config.Configurations.DBUser, config.Configurations.DBPassword,
			config.Configurations.DBName)
	} else if config.Configurations.DBProvider == "maria" {
		sqlDriver = "mysql"
		sqlconn = config.Configurations.DBUser + ":" + config.Configurations.DBPassword + "@tcp(" +
			config.Configurations.DBHost + ":" + config.Configurations.DBPort + ")/" + config.Configurations.DBName
	}

	db, err = sql.Open(sqlDriver, sqlconn)
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
	log.Println("fetching last record from " + config.Configurations.TableName + " table")
	query := `select entry, value from ` + config.Configurations.TableName + ` order by entry desc limit 1`
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
	status := rows.Next()
	if status {
		err = rows.Scan(&lastIndex, &lastValue)
		if err != nil {
			log.Println(`[ERROR occured] ` + err.Error())
			return 0, "", err
		}
		log.Println("Query response for last record entry = " + strconv.Itoa(lastIndex) + " value = " + lastValue)
	} else {
		log.Println("no elements are present in queue")
		lastIndex = 0
		lastValue = ""
		err = nil
	}

	return lastIndex, lastValue, err
}

func DeleteLastRecord(entry int) error {
	var err error
	deleteStatement := `delete from ` + config.Configurations.TableName + ` where entry=` + strconv.Itoa(entry)
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
	log.Println("inserting new record into " + config.Configurations.TableName + " table")
	valuesPlaceHolder := ""
	if config.Configurations.DBProvider == "postgres" {
		valuesPlaceHolder = "($1, $2)"
	} else if config.Configurations.DBProvider == "maria" {
		valuesPlaceHolder = "(?, ?)"
	}
	insertStatement := `insert into ` + config.Configurations.TableName + ` (entry, value) values ` + valuesPlaceHolder
	_, err = db.Exec(insertStatement, entry, value)
	if err != nil {
		log.Println(`[ERROR occured] ` + err.Error())
		return err
	}
	log.Println("insert successful")
	return err
}

func ReinitializeTable() error {
	var err error
	db, err := connectDB()
	if err != nil {
		log.Println(`[ERROR occured] ` + err.Error())
		return err
	}
	defer db.Close()
	log.Println("dropping table " + config.Configurations.TableName)
	dropStatement := ""
	if config.Configurations.DBProvider == "postgres" {
		dropStatement = `drop table "` + config.Configurations.TableName + `"`
	} else if config.Configurations.DBProvider == "maria" {
		dropStatement = `drop table if exists ` + config.Configurations.TableName
	}
	_, err = db.Exec(dropStatement)
	if err != nil {
		log.Println(`[ERROR occured] ` + err.Error())
		return err
	}
	log.Println(config.Configurations.TableName + " dropped successfully")
	createStatement := `create table ` + config.Configurations.TableName + ` (entry integer unique, value character varying(255))`
	_, err = db.Exec(createStatement)
	if err != nil {
		log.Println(`[ERROR occured] ` + err.Error())
		return err
	}
	return err
}
