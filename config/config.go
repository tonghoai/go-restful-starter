package conf

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	MYSQL_DB string
	MYSQL_GENERATOR_DB string
}

func GetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var config Config
	config.MYSQL_DB = os.Getenv("MYSQL_DB")
	config.MYSQL_GENERATOR_DB = os.Getenv("MYSQL_GENERATOR_DB")
	return config
}
