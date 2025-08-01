package entity

type Program struct {
	ID      int
	Name    string
	EngName string `gorm:"column:name_eng"`
}
