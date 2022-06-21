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

	// Configuration for Dev mode
	// used to allow user to send credential requests
	r.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"http://127.0.0.1:8080"},
			AllowMethods:     []string{"PUT", "GET", "POST"},
			AllowHeaders:     []string{"Access-Control-Allow-Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return origin == "https://127.0.0.1:8080"
			},
		},
	))
	r.GET("/", testPath)
	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", routes.Register)
		userRouter.GET("/register", routes.UserRegisterPage)
		userRouter.POST("/login", routes.Login)
		userRouter.GET("/login", routes.UserLoginPage)
		userRouter.GET("/check/:name", routes.UserCheck)
		userRouter.GET("/page", routes.MiddleAuth, routes.Page)
	}
	r.Run(":3000")
}

func testPath(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Good",
	})
}
