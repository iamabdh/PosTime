package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", testPath)
	r.Run(":3000")
}

func testPath(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Good",
	})
}
