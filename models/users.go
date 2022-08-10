package models

type User struct {
	Id          int64  `gorm:"primary_key autoIncrement"`
	Name        string `gorm:"size:50"`
	Surname     string `gorm:"size:50"`
	Mail        string `gorm:"size:100"`
	Password    string `gorm:"size:50"`
	PhoneNumber string `gorm:"size:25"`
	Country     string `gorm:"size:50"`
	City        string `gorm:"size:50"`
	Mac         string `gorm:"size:17"`
}
