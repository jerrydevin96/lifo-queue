package config

import (
	"log"
	"os"
)

//AppConfig contains all configurations necessary for the application
type AppConfig struct {
	DBHost            string
	DBPort            string
	DBUser            string
	DBPassword        string
	DBName            string
	TableName         string
	DBProvider        string
	ReinitializeTable string
}

var (
	//Configurations contain all runtime configs for use in other packages
	Configurations *AppConfig
)

//Configure reads environment variables and sets them to Configurations variable
func Configure() {
	log.Println("fetching configurations from environment")
	Configurations = new(AppConfig)
	Configurations.DBHost = os.Getenv("DBHOST")
	Configurations.DBPort = os.Getenv("DBPORT")
	Configurations.DBUser = os.Getenv("DBUSER")
	Configurations.DBPassword = os.Getenv("DBPASSWORD")
	Configurations.DBName = os.Getenv("DBNAME")
	Configurations.TableName = os.Getenv("TABLENAME")
	Configurations.DBProvider = os.Getenv("DBPROVIDER")
	Configurations.ReinitializeTable = os.Getenv("REINITIALIZE_TABLE")
	log.Println("DBHOST " + Configurations.DBHost)
	log.Println("DBPORT " + Configurations.DBPort)
	log.Println("DBUSER " + Configurations.DBUser)
	if len(Configurations.DBPassword) != 0 {
		log.Println("DBPASSWORD ********")
	} else {
		log.Println("DBPASSWORD is empty")
	}
	log.Println("DBNAME " + Configurations.DBName)
	log.Println("TABLENAME " + Configurations.TableName)
	log.Println("DBPROVIDER " + Configurations.DBProvider)
	log.Println("REINITIALIZE_TABLE " + Configurations.ReinitializeTable)
}
