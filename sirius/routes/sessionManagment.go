package routes

import (
	"PosTime/models"
	"crypto/rand"
	"encoding/hex"
	"fmt"
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
	ConnectionDB.Db.Create(models.Session{UID: UID, SID: token})
	return token
}

// Session ID User : used to get user id from session

func SessionIDUser(UID string) string {
	var userSessionAndID models.Session
	ConnectionDB.Db.Find(&userSessionAndID, "s_id = ?", UID)
	return userSessionAndID.UID
}
