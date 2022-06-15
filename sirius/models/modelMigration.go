package models

import "log"

func MigrateModel() {
	db := ConnectDatabase()
	if db.AutoMigrate(&User{}, &Session{}) != nil {
		log.Fatal("Cannot Migrate User")
	} else {
		log.Println("User Model Migrated Successfully")
	}
}
