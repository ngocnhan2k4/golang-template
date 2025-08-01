package entity

type Status struct {
	ID      int
	Name    string
	EngName string `gorm:"column:name_eng"`
	Order   int    `gorm:"column:status_order"`
}

func (s Status) TableName() string {
	return "student_statuses"
}
