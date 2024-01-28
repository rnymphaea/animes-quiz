package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Get(key string) (string, bool) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	return os.LookupEnv(key)
}
