package routes

import (
	b "PosTime/models"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// @desc Generate keys randomly
// future: update & secure
func _token() string {
	b := make([]byte, 10)
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
	token := _token()
	ConnectionDB.Db.Create(b.Session{UID: UID, SID: token})
	return token
}

// @desc Validate each request incoming from user

//func ValidateRequest(UID string) bool {
//
//}
