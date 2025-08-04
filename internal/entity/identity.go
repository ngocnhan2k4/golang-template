package entity

import "time"

type IdentityDocument struct {
	ID         int       `json:"-"`
	Type       string    `json:"type"`
	Number     string    `json:"documentNumber"`
	IssuedDate time.Time `json:"issueDate"`
	ExpiryDate time.Time `json:"expiryDate"`
	IssuePlace string    `json:"issuePlace"`
	Country    string    `json:"countryIssue"`
	IsChip     bool      `json:"isChip"`
	Note       string    `json:"notes"`
	StudentID  int       `json:"-" gorm:"foreignKey:StudentID;references:Id"`
}
