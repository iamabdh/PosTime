package routes

import "PosTime/models"

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

type MyPosTimer struct {
	PosTimeID string `json:"PosTimeId"`
	Text      string `json:"Text"`
	Time      string `json:"Time"`
}
