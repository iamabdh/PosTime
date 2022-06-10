package routes

import (
	"github.com/gin-gonic/gin"
	"log"
)

type UserRegisterData struct {
	Name     string
	Email    string
	Username string
	Password string
}

func Register(c *gin.Context) {
	var registerData UserRegisterData
	err := c.BindJSON(&registerData)
	if err != nil {
		log.Fatal("Register: Unable tp Parse Data")
	} else {
		//	Verify incoming user data

		c.JSON(200, gin.H{
			"status": "ok",
		})
	}
}

func Login() {

}
