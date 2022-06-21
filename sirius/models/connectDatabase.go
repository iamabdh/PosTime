package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnection struct {
	Db *gorm.DB
}

func (s *DBConnection) ConnectDatabase() {
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

	var err error
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)
	s.Db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		log.Fatal("Database Cannot Connected")
	} else {
		log.Println("Database Connected Successfully")
	}
}
