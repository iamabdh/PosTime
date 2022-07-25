package main

import (
	"PosTime/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
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
		userRouter.GET("/logout", routes.Logout)
		userRouter.GET("/login", routes.UserLoginPage)
		userRouter.GET("/check/:name", routes.UserCheck)
		userRouter.GET("/page", routes.MiddleAuth, routes.Page)
		userRouter.POST("/postime/create", routes.MiddleAuth, routes.CreatePosTime)
		userRouter.GET("/postime/my-postime", routes.MiddleAuth, routes.MyPosTime)
		userRouter.GET("/postime/my-postimer", routes.MiddleAuth, routes.MyPosTimer)
		userRouter.GET("/postime/public-postimers", routes.MiddleAuth, routes.PublicPostimers)
		userRouter.POST("/postime/new-postimer", routes.MiddleAuth, routes.UserNewPostimer)
		userRouter.GET("/postime/postimers", routes.MiddleAuth, routes.UserPostimers)
		userRouter.GET("/postime/feed-postimers", routes.MiddleAuth, routes.FeedPosTimers)
		userRouter.GET("/postime/low-profile", routes.MiddleAuth, routes.UserDataLowProfile)
		userRouter.GET("/postime/last-update", routes.MiddleAuth, routes.UserPosTimerLastUpdate)
		userRouter.GET("/postime/find-postimer", routes.MiddleAuth, routes.FindPosTimer)
	}
	r.Run(":3000")
}

func testPath(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Good",
	})
}
