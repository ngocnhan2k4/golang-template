package entity

import "time"

type Class struct {
	ID           int
	AcademicYear int
	CourseID     int
	Semester     int
	TeacherName  string
	MaxStudents  int
	Room         string
	DayOfWeek    int
	StartTime    float64
	EndTime      float64
	DeadLine     time.Time `gorm:"column:deadline"`
	IsDeleted    bool
	Course       Course `gorm:"foreignKey:CourseID;references:ID"`
}
