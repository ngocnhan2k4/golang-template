package entity

type Faculty struct {
	ID      int    `gorm:"primaryKey"`
	EngName string `gorm:"column:name_eng"`
	Name    string
	Courses []Course `gorm:"foreignKey:FacultyId;references:ID"`
}
