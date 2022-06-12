package main

import (
	b "PosTime/models"
	"PosTime/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	b.MigrateModel()
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", testPath)
	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", routes.Register)
		userRouter.GET("/register", routes.UserRegisterPage)
		userRouter.POST("/login", routes.Login)
		userRouter.GET("/login", routes.UserLoginPage)
		userRouter.GET("/page", routes.MiddleAuth, routes.Page)
	}
	r.Run(":3000")
}

func testPath(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Good",
	})
}
