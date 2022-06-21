package routes

import b "PosTime/models"

var ConnectionDB b.DBConnection

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
