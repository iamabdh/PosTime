package models

import "log"

func (s *DBConnection) MigrateModel() {
	if s.Db.AutoMigrate(&User{}, &Session{}, &PosTimersFriend{}, &PosTime{}, &LastUpdate{}) != nil {
		log.Fatal("Cannot Migrate User")
	} else {
		log.Println("User Model Migrated Successfully")
	}
}
