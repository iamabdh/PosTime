package models

import "time"

type User struct {
	ID           string `gorm:"default:uuid_generate_v4()"`
	Name         string `gorm:"not null;size:30"`
	Username     string `gorm:"not null;size:30"`
	Email        string `gorm:"not null; size:150"`
	Password     string `gorm:"not null;"`
	Gender       string
	DateOfBirth  string
	Bio          string
	ImageProfile string
	DateJoined   time.Time
}
