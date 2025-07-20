package entity

import "time"

type Student struct {
	Id               string `gorm:"type:varchar(8)"`
	Name             string `gorm:"type:varchar(50)"`
	DateOfBirth      time.Time
	Gender           Gender `gorm:"default:0"`
	Email            string `gorm:"type:varchar(50)"`
	Phone            string `gorm:"type:varchar(20)"`
	PermanentAddress string
	TemporaryAddress string
	MailingAddress   string
	Nationality      string `gorm:"type:varchar(50)"`
	Course           int
	Isdeleted        bool
}
