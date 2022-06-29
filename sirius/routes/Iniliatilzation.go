package routes

import (
	"PosTime/models"
	"fmt"
	"time"
)

var ConnectionDB models.DBConnection

func init() {
	ConnectionDB.ConnectDatabase()
	ConnectionDB.MigrateModel()
}

type UserRegisterData struct {
	Name     string
	Email    string
	Username string
	Password string
}

type UserLoginData struct {
	Username string
	Password string
}

type UserLoggedData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserPosTimeCreate struct {
	Text string
}

type UserSession struct {
	SessionID string
	UserID    string
}

type PosTime struct {
	PosTimeID string `json:"PosTimeId"`
	Text      string `json:"Text"`
	Time      string `json:"Time"`
	Date      string `json:"Date"`
}

type PublicPostimerProfile struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

type NewPosTimer struct {
	Username string
}

func CallDate() string {
	return fmt.Sprintf("%02d-%02d-%d",
		time.Now().Day(),
		time.Now().Month(),
		time.Now().Year())
}

func CallTime() string {
	return fmt.Sprintf("%02d:%02d",
		time.Now().Hour(),
		time.Now().Minute())
}
