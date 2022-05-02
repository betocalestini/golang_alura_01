package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Use(key string) string {
	var retorno string
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	database := os.Getenv("DATABASE")
	db_user := os.Getenv("DB_USER")
	password := os.Getenv("PASSWORD")
	db_port := os.Getenv("DB_PORT")

	switch key {
	case "HOST":
		retorno = host
	case "PORT":
		retorno = port
	case "DB_PORT":
		retorno = db_port
	case "DATABASE":
		retorno = database
	case "DB_USER":
		retorno = db_user
	case "PASSWORD":
		retorno = password
	}

	return retorno
}
