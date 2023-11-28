package main

import (
	"users_manager/internal/migration"
	"users_manager/internal/services/database"
	"users_manager/internal/services/http"

	"github.com/Valgard/godotenv"
)

func init() {
	dotenv := godotenv.New()
	if err := dotenv.Load(".env"); err != nil {
		panic(err)
	}
	database.Connection()
	migration.Mingrations()
	http.HttpClient()
}

func main() {

}
