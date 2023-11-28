package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Conn *gorm.DB
	err  error
)

func Connection() {

	connStr := fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	Conn, err = gorm.Open(postgres.Open(connStr))

	if err != nil {
		fmt.Println("ERR: ", err)
		log.Fatal("POSTGRESQL CONNECTION FAILED")
	} else {
		fmt.Println("POSTGRESQL CONNECTION SUCCESS")
	}
}
