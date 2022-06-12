package routes

import (
	b "PosTime/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
)

var store = sessions.NewCookieStore([]byte("test"))

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
		return
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
	// Check password if it's correct
	if !(bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)) == nil) {
		c.JSON(401, gin.H{
			"status":  "Error: Wrong Password",
			"forward": "/login",
		})
		return
	}

	// Set cookie to user for new login
	session, _ := store.Get(c.Request, "session")
	session.Values["user"] = user.Username
	sessionErr := session.Save(c.Request, c.Writer)
	// Check error if it exists
	if sessionErr != nil {
		log.Fatal(sessionErr)
		c.JSON(406, gin.H{
			"status": "bad",
		})
		return
	}
	// Redirect user to page profile
	c.Redirect(301, "/user/page")
}

func MiddleAuth(c *gin.Context) {
	fmt.Println("running")
	session, _ := store.Get(c.Request, "session")
	_, ok := session.Values["user"]
	if !ok {
		c.Redirect(301, "/user/login")
		c.Abort()
		return
	}
	c.Next()
}

func Page(c *gin.Context) {
	//cookie, err := c.Request.Cookie("session")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//cookie.Value
	session, _ := store.Get(c.Request, "session")
	val := session.Values["user"]
	var user b.User
	// Request: to database for this username
	b.ConnectDatabase().Find(&user, "username = ?", val)
	// Send profile data
	loggedUser := UserLoggedData{Name: user.Name, Email: user.Email, Username: user.Username}
	loggedUserJson, _ := json.Marshal(loggedUser)
	c.JSON(200, string(loggedUserJson))
}
