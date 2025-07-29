package entity

import "time"


type Course struct{
	ID int
	EngName string
	Name string
	Credits int
	FacultyId int
	Description string
	DescriptionEng string
	RequiredCourseId string
	DeletedAt time.Time
	CreatedAt time.Time
}