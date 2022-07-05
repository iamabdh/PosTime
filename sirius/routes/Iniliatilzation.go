package routes

import (
	"PosTime/models"
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
	PosTimeID string    `json:"PosTimeId"`
	Username  string    `json:"Username"`
	Text      string    `json:"Text"`
	Date      time.Time `json:"Date"`
}

type PublicPostimerProfile struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

type NewPosTimer struct {
	Username string
}

type DataLowProfile struct {
	Name       string `json:"name"`
	Username   string `json:"username"`
	Postime    int64  `json:"postime"`
	Postimer   int64  `json:"postimer"`
	LastUpdate string `json:"lastUpdate"`
}
