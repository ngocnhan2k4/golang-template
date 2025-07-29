package entity

import "time"


type Class struct{
	ID int
	AcademicYear int
	CourseID int
	Semester int
	TeacherName string
	MaxStudents int
	Room string
	DayOfWeek int
	StartTime time.Time
	EndTime time.Time
	DeadLine time.Time
	IsDeleted bool 
}