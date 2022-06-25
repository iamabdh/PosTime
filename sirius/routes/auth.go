package routes

import (
	"PosTime/models"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// User login and Register
// @route GET /user/login

func UserLoginPage(c *gin.Context) {
}

// @route GET /user/register

func UserRegisterPage(c *gin.Context) {
}

// User register & login
// @route POST /user/register

func Register(c *gin.Context) {
	var registerData UserRegisterData
	err := c.BindJSON(&registerData)
	if err != nil {
		fmt.Println("Register, POST: Unable to Parse Data", err)
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

			ConnectionDB.Db.Create(&models.User{
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
		return
	}
}

// Login @route POST /user/login

func Login(c *gin.Context) {
	var loginData UserLoginData
	err := c.BindJSON(&loginData)
	if err != nil {
		fmt.Println("Error: Unable to Parse Login Data")
		c.JSON(400, gin.H{
			"status": "Error: Unable to Parse Login Data",
		})
		return
	}
	var user models.User
	// Request: to database for this username
	ConnectionDB.Db.Find(&user, "username = ?", loginData.Username)
	// Check password if it's correct
	if !(bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)) == nil) {
		c.JSON(401, gin.H{
			"status":  "Error: Wrong Password",
			"forward": "/user/login",
		})
		return
	}

	// Set cookie to user for new login
	// Save user session id on sessions table
	token := ValidateCookie(user.ID)
	c.SetCookie("session", token, 400, "/", "http://127.0.0.1:3000", false, true)
	c.JSON(200, gin.H{
		"status": "logged successfully",
		"allow":  true,
	})
}

// Logout user
// middleware required
func Logout(c *gin.Context) {

}

// Auth middleware for checking auth of user

func MiddleAuth(c *gin.Context) {
	// find the cookie from request
	// require to check issue data of time
	session, err := c.Cookie("session")
	var sessionID models.Session
	ConnectionDB.Db.Find(&sessionID, "s_id = ?", session)
	if err != nil || sessionID.SID == "" {
		fmt.Println("Middle Auth: ", err)
		c.JSON(401, gin.H{
			"status":  "wrong",
			"forward": "login",
			"allow":   false,
		})
		c.Abort()
		return
	}
	c.Next()
}

// check user auth before each request
// this prevents accessing pages such as register & login
// @route GET /user/check/:id

func UserCheck(c *gin.Context) {
	name := c.Params[0].Value
	session, err := c.Cookie("session")
	var sessionID models.Session
	ConnectionDB.Db.Find(&sessionID, "s_id = ?", session)
	fmt.Println("err: ", err)
	fmt.Println("session: ", session)
	fmt.Println("name: ", name)
	if (err != nil || sessionID.SID == "") && name != "page" {
		c.JSON(200, gin.H{
			"forward": name,
		})
		return
	} else if (err != nil || sessionID.SID == "") && name == "page" {
		c.JSON(200, gin.H{
			"forward": "login",
		})
		return
	}
	switch name {
	case "login":
		c.JSON(200, gin.H{
			"forward": "/",
		})
		break
	case "register":
		c.JSON(200, gin.H{
			"forward": "/",
		})
		break
	case "page":
		c.JSON(200, gin.H{
			"forward": "/",
		})
	}
}

// @route GET /user/page

func Page(c *gin.Context) {
	session, err := c.Cookie("session")
	if err != nil {
		fmt.Println("Page, GET: ", err)
		c.JSON(200, gin.H{
			"status":  "something wrong",
			"forward": "/user/login",
		})
		return
	}
	// get user id from user session
	var sessionID models.Session
	ConnectionDB.Db.Find(&sessionID, "s_id = ?", session)
	var user models.User
	// Request: to database for this username
	ConnectionDB.Db.Find(&user, "id = ?", sessionID.UID)
	// Send profile data
	loggedUser := UserLoggedData{
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username}
	loggedUserJson, _ := json.Marshal(loggedUser)
	c.JSON(200, string(loggedUserJson))
}
