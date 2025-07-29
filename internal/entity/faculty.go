package entity




type Faculty struct{
	ID int 
	EngName string `gorm:"column:name_eng"`
	Name string 
}