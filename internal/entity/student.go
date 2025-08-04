package entity

import "time"

type Student struct {
	Id               int
	Name             string
	DateOfBirth      time.Time
	Gender           string
	Email            string
	Course           int
	Phone            string
	PermanentAddress string
	TemporaryAddress string
	MailingAddress   string
	ProgramID        int
	StatusID         int
	FacultyID        int
	Nationality      string
	IsDeleted        bool
}
