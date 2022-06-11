package routes

import (
	b "PosTime/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
)

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

func Register(c *gin.Context) {
	var registerData UserRegisterData
	err := c.BindJSON(&registerData)
	if err != nil {
		log.Fatal("Register: Unable to Parse Data")
		// Send to user the err with code
		c.JSON(406, gin.H{
			"status": "Error: Unable to Parse Data",
		})
	} else {
		// Verify incoming user data
		// using regex for all string patterns
		// compile all required patterns

		rName, _ := regexp.Compile("^[a-zA-Z]+$")
		rEmail, _ := regexp.Compile("(\\w\\.?)+@[\\w\\.-]+\\.\\w{2,}")
		rUsername, _ := regexp.Compile("[a-zA-Z0-9-_]+${4,}")

		// name & email & username
		if !(rName.MatchString(registerData.Name) && rEmail.MatchString(registerData.Email) && rUsername.MatchString(registerData.Username)) {
			c.JSON(406, gin.H{
				"status": "issue",
			})
			return
		} else {
			// Check unique email & username ...
			// Encrypt user password
			hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(registerData.Password), 0)
			// Store all data to database
			b.ConnectDatabase().Create(&b.User{
				Name:     registerData.Name,
				Email:    registerData.Email,
				Username: registerData.Username,
				Password: string(hashedPassword),
			})
			c.JSON(200, gin.H{
				"status":  "ok",
				"forward": "/login",
			})
		}
	}
}

func Login(c *gin.Context) {
	var loginData UserLoginData
	err := c.BindJSON(&loginData)
	if err != nil {
		log.Fatal("Error: Unable to Parse Login Data")
		c.JSON(400, gin.H{
			"status": "Error: Unable to Parse Login Data",
		})
		return
	}
	var user b.User
	// Request: to database for this username
	b.ConnectDatabase().Find(&user, "username = ?", loginData.Username)
	fmt.Println(user)
	// Check password if it's correct
	if !(bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)) == nil) {
		c.JSON(401, gin.H{
			"status":  "Error: Wrong Password",
			"forward": "/login",
		})
	}
	loggedUser := UserLoggedData{Name: user.Name, Email: user.Email, Username: user.Username}
	loggedUserJson, _ := json.Marshal(loggedUser)
	c.JSON(200,
		string(loggedUserJson),
	)
}
