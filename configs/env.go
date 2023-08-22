package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvGetURI(env string) string {
	err := godotenv.Load("./../.env")

	if err != nil {
		log.Fatalln("error .env")
	}

	data := os.Getenv(env)
	return data
}
