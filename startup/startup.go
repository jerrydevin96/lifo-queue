package startup

import (
	"log"

	"github.com/jerrydevin96/lifo-queue/config"
	"github.com/jerrydevin96/lifo-queue/database"
)

//AppStartup initializes the application
func AppStartup() error {
	var err error
	log.Println("Initializing application")
	config.Configure()
	if config.Configurations.ReinitializeTable == "Y" {
		log.Println("Reinitializing table " + config.Configurations.TableName)
		err = database.ReinitializeTable()
		if err != nil {
			log.Println("[ERROR] error occured " + err.Error())
			return err
		}
		log.Println("database initialization completed")
	}

	return err
}
