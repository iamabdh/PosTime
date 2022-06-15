package routes

import (
	b "PosTime/models"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

// User login and Register
// @GET /user/login

func UserLoginPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
		"page":   "Login Page",
	})
}

// @GET /user/register

func UserRegisterPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
		"page":   "Register Page",
	})
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
		return
	}
}

// @route POST /user/login
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
	var user b.User
	// Request: to database for this username
	b.ConnectDatabase().Find(&user, "username = ?", loginData.Username)
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
	c.SetCookie("session", token, 100, "/", "http://127.0.0.1:3000", false, true)
	c.Redirect(301, "/user/page")
}

// Logout user
// middleware required
func Logout(c *gin.Context) {

}

// Auth middleware for checking auth of user

func MiddleAuth(c *gin.Context) {

	//var user b.User
	// find the cookie from request
	// require to check issue data of time
	if _, err := c.Cookie("session"); err != nil {
		c.JSON(200, gin.H{
			"status":  "wrong",
			"forward": "/user/login",
		})
		c.Abort()
		return
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
	var sessionID b.Session
	b.ConnectDatabase().Find(&sessionID, "s_id = ?", session)
	var user b.User
	// Request: to database for this username
	b.ConnectDatabase().Find(&user, "id = ?", sessionID.UID)
	// Send profile data
	loggedUser := UserLoggedData{
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username}
	loggedUserJson, _ := json.Marshal(loggedUser)
	c.JSON(200, string(loggedUserJson))
}
