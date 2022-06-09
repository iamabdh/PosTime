package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	if godotenv.Load("config/.env") != nil {
		log.Fatal("Unable to Load Env File")
	} else {
		log.Println("Env File Loaded Successfully")
	}

	// Load env vars

	user := os.Getenv("user")
	port := os.Getenv("port")
	host := os.Getenv("host")
	dbName := os.Getenv("dbName")
	password := os.Getenv("password")

	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		log.Fatal("Database Cannot Connected")
	} else {
		log.Println("Database Connected Successfully")
	}
	return db
}
