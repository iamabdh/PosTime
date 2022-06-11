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
		userRouter.POST("/login", routes.Login)
	}
	r.Run(":3000")
}

func testPath(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Good",
	})
}
