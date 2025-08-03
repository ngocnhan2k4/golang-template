package entity

import "time"

type Course struct {
	ID               int
	EngName          string `gorm:"column:name_eng"`
	Name             string
	Credits          int
	FacultyId        int
	Description      string
	DescriptionEng   string
	RequiredCourseId *int
	DeletedAt        time.Time
	CreatedAt        time.Time
	Faculty          Faculty `gorm:"foreignKey:FacultyId;references:ID"`
}
