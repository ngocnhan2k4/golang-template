package entity

import "time"

type Student struct {
	Id               string    `gorm:"type:varchar(8)" json:"id" example:"22120249" format:"string"`
	Name             string    `gorm:"type:varchar(50)" json:"name" example:"Tran Ngoc Nhan" format:"string"`
	DateOfBirth      time.Time `json:"date_of_birth" example:"2004-07-05" format:"date"`
	Gender           Gender    `gorm:"default:0" json:"gender" example:"0" format:"int"`
	Email            string    `gorm:"type:varchar(50)" json:"email" example:"abc@gmail.com" format:"string"`
	Phone            string    `gorm:"type:varchar(20)" json:"phone" example:"0123456789" format:"string"`
	PermanentAddress string    `json:"permanent_address" example:"123 Main St, City, Country" format:"string"`
	TemporaryAddress string    `json:"temporary_address" example:"456 Elm St, City, Country" format:"string"`
	MailingAddress   string    `json:"mailing_address" example:"789 Oak St, City, Country" format:"string"`
	Nationality      string    `gorm:"type:varchar(50)" json:"nationality" example:"Vietnamese" format:"string"`
	Course           int       `json:"course" example:"2022" format:"int"`
	Isdeleted        bool      `json:"isdeleted" example:"false" format:"bool"`
}
