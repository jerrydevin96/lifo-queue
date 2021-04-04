package config

import (
	"os"
	"testing"
)

func SetEnvironment(dbHost string, dbPort string, dbUser string, dbPassword string,
	dbName string, dbTable string, dbProvider string, reinitialize string) {
	os.Setenv("DBHOST", dbHost)
	os.Setenv("DBPORT", dbPort)
	os.Setenv("DBUSER", dbUser)
	os.Setenv("DBPASSWORD", dbPassword)
	os.Setenv("DBNAME", dbName)
	os.Setenv("TABLENAME", dbTable)
	os.Setenv("DBPROVIDER", dbProvider)
	os.Setenv("REINITIALIZE_TABLE", reinitialize)
}

func validateConfig(dbHost string, dbPort string, dbUser string, dbPassword string,
	dbName string, dbTable string, dbProvider string, reinitialize string) bool {
	if Configurations.DBHost != dbHost {
		return false
	}
	if Configurations.DBPort != dbPort {
		return false
	}
	if Configurations.DBUser != dbUser {
		return false
	}
	if Configurations.DBPassword != dbPassword {
		return false
	}
	if Configurations.DBName != dbName {
		return false
	}
	if Configurations.TableName != dbTable {
		return false
	}
	if Configurations.DBProvider != dbProvider {
		return false
	}
	if Configurations.ReinitializeTable != reinitialize {
		return false
	}
	return true
}

func TestConfigure(t *testing.T) {
	SetEnvironment("10.11.20.30", "5432", "admin", "admin", "lifo", "lifo", "postgres", "N")
	Configure()
	validation := validateConfig("10.11.20.30", "5432", "admin", "admin", "lifo", "lifo", "postgres", "N")
	if !validation {
		t.Log("Validation failed")
		t.FailNow()
	} else {
		t.Log("Validation successful")
	}
}
