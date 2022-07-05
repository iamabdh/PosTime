package routes

import (
	"PosTime/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// CreatePosTime @desc POST /user/postime/create
func CreatePosTime(c *gin.Context) {
	_userSessionID, err := c.Cookie("session")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "wrong",
		})
		return
	}
	var newPostime UserPosTimeCreate
	if c.BindJSON(&newPostime) != nil || strings.TrimSpace(newPostime.Text) == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": "wrong postime",
		})
		return
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

// DestroyPosTime @desc DELETE /user/postime/destroy
func DestroyPosTime(c *gin.Context) {

}

// MyPosTimes fetch to a user his/her postimes
// @route GET /user/postime/my-postime
func MyPosTimes(c *gin.Context) {
	_userSessionID, err := c.Cookie("session")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "wrong",
		})
		return
	}
	_userID := SessionIDUser(_userSessionID)
	var userPosTimers []PosTime
	ConnectionDB.Db.Table("users").Joins("INNER JOIN pos_times on users.id=pos_times.source_pos_timer_id").Where("id = ?", _userID).Select("pos_time_id, username, text, time, date").Find(&userPosTimers)
	c.JSON(200, userPosTimers)
}

// PublicPostimers @route	GET /user/postime/public-postimers
// @desc	fetch to public all users
// {{API deprecate}}
func PublicPostimers(c *gin.Context) {
	var postimers []PublicPostimerProfile
	ConnectionDB.Db.Table("users").Select("name, username").Find(&postimers)
	c.JSON(http.StatusAccepted, postimers)
}

// UserNewPostimer
// @route	POST /user/postime/new-postimer
// @desc	user add new postimer
func UserNewPostimer(c *gin.Context) {
	_userSessionID, err := c.Cookie("session")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "wrong",
			"error":  "cannot parse session id",
		})
		return
	}
	var username NewPosTimer
	if c.BindJSON(&username) != nil || strings.TrimSpace(username.Username) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "wrong",
			"error":  "username not exist",
		})
		return
	}
	_userID := SessionIDUser(_userSessionID)
	ok, _targetUserID := IDUserLookUp(username.Username)
	if ok || _targetUserID == _userID {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":   "wrong",
			"username": "not exist or duplication occur",
		})
		return
	}
	// check user if the already has friendship with this postimer username
	if FriendshipLookUp(_userID, _targetUserID) {
		ConnectionDB.Db.Create(&models.PosTimersFriend{
			SourceFriendID: _userID,
			TargetFriendID: _targetUserID,
		})
	}
	c.JSON(200, gin.H{
		"status":  "good",
		"message": "Friendship added successfully",
	})
}

// UserPostimers
// @route	GET /user/postime/postimers
// @desc	fetch to user all his/her postimers ==> friends
func UserPostimers(c *gin.Context) {
	_userSessionID, err := c.Cookie("session")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "wrong",
			"error":  "cannot parse session id",
		})
		return
	}
	_userID := SessionIDUser(_userSessionID)
	var userPostimer []string
	ConnectionDB.Db.Table("users").Joins(
		"INNER JOIN pos_timers_friends on users.id=pos_timers_friends.source_friend_id").Where(
		"source_friend_id = ?", _userID).Select(
		"target_friend_id").Find(&userPostimer)

	// Get all user postimers & fed to json array
	var userPostimersAll []PublicPostimerProfile
	for _, valID := range userPostimer {
		var userPostimerData PublicPostimerProfile
		ConnectionDB.Db.Table("users").Where("id = ?", valID).Select("name, username").Find(&userPostimerData)
		userPostimersAll = append(userPostimersAll, userPostimerData)
	}
	c.JSON(http.StatusAccepted, userPostimersAll)
}

// UserRemovePostimer remove postimer from user
// route DELETE /user/postime/remove-postime
func UserRemovePostimer(c *gin.Context) {

}

// FeedPosTimers  fetch to users all postimers in user feed
// followed by time
// future: followed by priority
// @route GET /user/postime/feed-postimers
func FeedPosTimers(c *gin.Context) {
	_userSessionID, err := c.Cookie("session")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "wrong",
		})
		return
	}
	// Get all user friends ids & userid
	_usersIDs := FriendShipIDs(SessionIDUser(_userSessionID))
	fmt.Println(_usersIDs)
	var userPosTimersAll []PosTime
	for _, valID := range _usersIDs {
		var userPosTimers []PosTime
		ConnectionDB.Db.Table("users").Joins("INNER JOIN pos_times on users.id=pos_times.source_pos_timer_id").
			Where("id = ?", valID).
			Select("pos_time_id, username, text, date").
			Find(&userPosTimers)
		//fmt.Println(userPosTimers)
		userPosTimersAll = append(userPosTimersAll, userPosTimers...)
	}
	c.JSON(http.StatusAccepted, userPosTimersAll)
}

// UserDataLowProfile
// @route GET /user/postime/low-profile
// @desc data required for main page
// redundant request may be changed or deleted in future !
func UserDataLowProfile(c *gin.Context) {
	_userSessionID, err := c.Cookie("session")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "wrong",
		})
		return
	}
	_userID := SessionIDUser(_userSessionID)
	var userLowProfile DataLowProfile

	// assign name and username
	ConnectionDB.Db.Table("users").
		Where("id = ?", _userID).
		Select("name, username").Find(&userLowProfile)

	// assign number of postimes and postimer
	ConnectionDB.Db.Table("users").
		Joins("INNER JOIN pos_times on users.id=pos_times.source_pos_timer_id").
		Where("id = ? ", _userID).
		Count(&userLowProfile.Postime)

	ConnectionDB.Db.Table("users").
		Joins("INNER JOIN pos_timers_friends on users.id=pos_timers_friends.source_friend_id").
		Where("id = ? ", _userID).
		Count(&userLowProfile.Postimer)
	c.JSON(http.StatusAccepted, userLowProfile)
}

// FindPosTimer
// @route GET /user/postime/find-postimer
// @desc send to user collection of postimers (not followed)
func FindPosTimer(c *gin.Context) {
	_userSessionID, err := c.Cookie("session")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "wrong",
		})
		return
	}
	_userID := SessionIDUser(_userSessionID)
	var usersLowProfile []DataLowProfile
	// find all user that not followed
	ConnectionDB.Db.Table("users").
		Joins("RIGHT JOIN pos_timers_friends on users.id=pos_timers_friends.source_friend_id").
		Where("id = ? ", _userID).
		Select("name, username").Find(&usersLowProfile)

	c.JSON(http.StatusAccepted, usersLowProfile)
}
