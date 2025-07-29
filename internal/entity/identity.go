package entity

import "time"


type IdentityDocument struct{
	ID int
	Type string
	Number string
	IssuedDate time.Time
	ExpiryDate time.Time
	IssuePlace string
	Country string
	IsChip bool
	Note string
	StudentID int
}