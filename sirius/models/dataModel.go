package models

type User struct {
	ID           string `gorm:"default:uuid_generate_v4()"`
	Name         string `gorm:"not null;size:30"`
	Username     string `gorm:"not null;size:30"`
	Email        string `gorm:"not null; size:150"`
	Password     string `gorm:"not null;"`
	Gender       string
	DateOfBirth  string
	Bio          string
	ImageProfile string
	DateJoined   string `gorm:"not null"`
}

type Session struct {
	UID  string `gorm:"not null"`
	SID  string `gorm:"not null"`
	Time string `gorm:"not null"`
	Date string `gorm:"not null"`
}

type PosTimersFriend struct {
	SourceFriendID string
	TargetFriendID string
	User1          User `gorm:"foreignKey:SourceFriendID;references:ID"`
	User2          User `gorm:"foreignKey:TargetFriendID;references:ID"`
}

type PosTime struct {
	PosTimeID        string `gorm:"not null"`
	SourcePosTimerID string `gorm:"not null"`
	Text             string `gorm:"not null"`
	Time             string `gorm:"not null"`
	Date             string `gorm:"not null"`
	User             User   `gorm:"foreignKey:SourcePosTimerID;references:ID"`
}
