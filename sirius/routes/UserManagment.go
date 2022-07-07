package routes

import (
	"PosTime/models"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

// @desc Generate keys randomly
// future: update & secure

func Token(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		fmt.Println(err)
		return ""
	}
	return hex.EncodeToString(b)
}

/*
@desc store session id
@param UID user id
for now only it stores ID
this func for future
*/

func _storeSession(UID string) {

}

func ValidateCookie(UID string) string {
	// Generate key for user
	token := Token(10)
	ConnectionDB.Db.Create(&models.Session{
		UID: UID,
		SID: token,
	})
	return token
}

// SessionIDUser Session ID User : used to get user id from session
func SessionIDUser(UID string) string {
	var userSessionAndID models.Session
	ConnectionDB.Db.Find(&userSessionAndID, "s_id = ?", UID)
	return userSessionAndID.UID
}

// IDUserLookUp Used to check username look-up & return id of user
func IDUserLookUp(username string) (bool, string) {
	var _userID string
	ConnectionDB.Db.Table("users").Where("username = ?", username).Select("id").Find(&_userID)
	return _userID == "", _userID
}

// FriendshipLookUp Check friendship if not existed added
// @params:
func FriendshipLookUp(sourceID, targetID string) bool {
	var _userIDTarget string
	ConnectionDB.Db.Table("users").Joins("INNER JOIN pos_timers_friends on users.id=pos_timers_friends.source_friend_id").Where("target_friend_id = ?", targetID).Select("target_friend_id").Find(&_userIDTarget)
	if _userIDTarget == "" {
		return true
	}
	return false
}

// FriendShipIDs
// @desc used for return list of friend id
// @params: addUser check if needed to add user id to list
// addUser ==> true for feed
func FriendShipIDs(addUser bool, _userId string) []string {
	var userPostimer []string
	ConnectionDB.Db.Table("users").Joins(
		"INNER JOIN pos_timers_friends on users.id=pos_timers_friends.source_friend_id").Where(
		"source_friend_id = ?", _userId).Select(
		"target_friend_id").Find(&userPostimer)
	if addUser {
		userPostimer = append(userPostimer, _userId)
	}
	return userPostimer
}

// GetPosTimeCreatedAt
// @desc Used for retrieve data the postime create
func GetPosTimeCreatedAt(postimeID string) time.Time {
	var _createAt time.Time
	ConnectionDB.Db.Table("pos_times").
		Where("pos_time_id = ?", postimeID).
		Select("date").
		Find(&_createAt)
	return _createAt
}