package routes

import (
	"PosTime/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// @desc POST /user/postime/create

func CreatePosTime(c *gin.Context) {
	var newPostime UserPosTimeCreate
	if c.BindJSON(&newPostime) != nil || strings.TrimSpace(newPostime.Text) == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": "wrong postime",
		})
		return
	}
	_userSessionID, err := c.Cookie("session")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "wrong",
		})
	}
	_userID := SessionIDUser(_userSessionID)
	ConnectionDB.Db.Create(&models.PosTime{
		PosTimeID:        Token(8),
		SourcePosTimerID: _userID,
		Text:             newPostime.Text,
	})
	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "PosTime Created Successfully",
	})
}

// @desc DELETE /user/postime/destroy

func DestroyPosTime(c *gin.Context) {

}

// fetch to a user his/her postimes
// @route GET /user/postime/mypostime

func MyPosTimers(c *gin.Context) {

}

// fetch to users all postimers in user feed
// followed by time
// future: followed by priority
// @desc GET /user/postime/postimers

func PosTimers(c *gin.Context) {

}
