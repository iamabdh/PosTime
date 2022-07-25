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
	_postimeToken := Token(8)

	ConnectionDB.Db.Create(&models.PosTime{
		PosTimeID:        _postimeToken,
		SourcePosTimerID: _userID,
		Text:             newPostime.Text,
	})

	// update last_update tables for this creation
	var _userIDLastUpdate string
	ConnectionDB.Db.Table("last_updates").
		Where("source_pos_timer_id = ?", _userID).
		Select("source_pos_timer_id").
		Find(&_userIDLastUpdate)

	if _userIDLastUpdate == "" {
		fmt.Println("first condition")
		ConnectionDB.Db.Create(&models.LastUpdate{
			PosTimeIDCreated: _postimeToken,
			SourcePosTimerID: _userID,
		})
	} else if _userIDLastUpdate == _userID {
		fmt.Println("second condition")
		ConnectionDB.Db.Table("last_updates").
			Where("source_pos_timer_id = ?", _userID).
			Updates(models.LastUpdate{
				PosTimeIDCreated: _postimeToken,
				Date:             GetPosTimeCreatedAt(_postimeToken),
			})
	}

	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "PosTime Created Successfully",
	})
}

// DestroyPosTime @desc DELETE /user/postime/destroy
func DestroyPosTime(c *gin.Context) {

}

// MyPosTime fetch to a user his/her postimes
// @route GET /user/postime/my-postime
func MyPosTime(c *gin.Context) {
	_userSessionID, err := c.Cookie("session")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "wrong",
		})
		return
	}
	_userID := SessionIDUser(_userSessionID)
	var userPosTime []PosTime
	ConnectionDB.Db.Table("users").
		Joins("INNER JOIN pos_times on users.id=pos_times.source_pos_timer_id").
		Where("id = ?", _userID).
		Select("pos_time_id, username, text, time, date").Find(&userPosTime)
	c.JSON(200, userPosTime)
}

// MyPosTimer fetch to a user postimer aka friends
// @route GET /user/postime/my-postimer
func MyPosTimer(c *gin.Context) {
	_userSessionID, err := c.Cookie("session")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "wrong",
		})
		return
	}
	_userID := SessionIDUser(_userSessionID)
	var userPosTimer []UserPosTimer
	ConnectionDB.Db.Table("users u1").
		Joins("INNER JOIN pos_timers_friends ptf on u1.id=ptf.source_friend_id").
		Joins("INNER JOIN users u2 on ptf.target_friend_id=u2.id").
		Where("u1.id = ?", _userID).
		Select("u2.name, u2.username").
		Find(&userPosTimer)
	c.JSON(200, userPosTimer)
}

// PublicPostimers @route	GET /user/postime/public-postimers
// @desc	fetch to public all users
// API DEV
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
			"check":    false,
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
		"check":   true,
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
	_usersID := SessionIDUser(_userSessionID)
	var feedUser []PosTime
	raw := fmt.Sprintf("select pt.pos_time_id, u2.username, pt.text, pt.date\n"+
		"from pos_timers_friends p1\n"+
		"inner join users u1 on u1.id = p1.source_friend_id\n"+
		"inner join users u2 on u2.id = p1.target_friend_id\n"+
		"inner join pos_times pt on u2.id = pt.source_pos_timer_id\n"+
		"where u1.id = '%[1]v'\n"+
		"union all\n"+
		"select pt.pos_time_id, u3.username, pt.text, pt.date\n"+
		"from users u3 inner join pos_times pt on u3.id = pt.source_pos_timer_id\n"+
		"where u3.id = '%[1]v'\n"+
		"order by date desc", _usersID)
	ConnectionDB.Db.Raw(raw).Find(&feedUser)
	c.JSON(http.StatusAccepted, feedUser)
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

// UserPosTimerLastUpdate
// @route GET /user/postime/last-update
// @desc Fetch to user last update from postimer
func UserPosTimerLastUpdate(c *gin.Context) {
	_userSessionID, err := c.Cookie("session")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "wrong",
		})
		return
	}
	var usersLastUpdate []PosTimeLastUpdate
	ConnectionDB.Db.Table("pos_timers_friends").
		Joins("inner join users on users.id = pos_timers_friends.target_friend_id").
		Joins("inner join last_updates on last_updates.source_pos_timer_id = users.id").
		Select("pos_time_id_created, username, date").
		Where("pos_timers_friends.source_friend_id = ?", SessionIDUser(_userSessionID)).
		Order("date desc").
		Find(&usersLastUpdate)
	c.JSON(http.StatusAccepted, usersLastUpdate)
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
	raw := fmt.Sprintf(
		"select u1.name, u1.username, lu.date\n"+
			"from users u1\n"+
			"left join last_updates lu on u1.id = lu.source_pos_timer_id\n"+
			"where not u1.id in (\n"+
			"    select f.target_friend_id\n"+
			"    from pos_timers_friends f\n"+
			"    where u1.id = f.target_friend_id and f.source_friend_id = '%[1]v'\n"+
			")\n"+
			"and u1.id != '%[1]v'\n", _userID)
	ConnectionDB.Db.Raw(raw).Find(&usersLowProfile)
	for index, valIDUsers := range usersLowProfile {
		// count postimes and postimers
		ConnectionDB.Db.Table("users").
			Joins("INNER JOIN pos_times on users.id=pos_times.source_pos_timer_id").
			Where("username = ? ", valIDUsers.Username).
			Count(&usersLowProfile[index].Postime)
		ConnectionDB.Db.Table("users").
			Joins("INNER JOIN pos_timers_friends on users.id=pos_timers_friends.source_friend_id").
			Where("username = ? ", valIDUsers.Username).
			Count(&usersLowProfile[index].Postimer)
	}
	c.JSON(http.StatusAccepted, usersLowProfile)
}
