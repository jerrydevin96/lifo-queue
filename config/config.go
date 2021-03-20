package config

//AppConfig contains all configurations necessary for the application
type AppConfig struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBProvider string
}

var (
	//Configurations contain all runtime configs for use in other packages
	Configurations *AppConfig
)

func configure() {
	Configurations = new(AppConfig)
}
