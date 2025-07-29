package entity

import "time"



type RegisterCancellationHistory struct{
	ID int
	ClassId int
	CourseName string
	StudentId int
	StudentName string
	Semester int
	AcademicYear int
	Time time.Time
}