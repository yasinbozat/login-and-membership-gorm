package models

import (
	"crypto/md5"
	"db_connect_app/database"
	"encoding/hex"
	"fmt"
)

type User struct {
	Id          int64  `gorm:"primary_key autoIncrement"`
	Name        string `gorm:"size:50"`
	Surname     string `gorm:"size:50"`
	Mail        string `gorm:"size:100"`
	Password    string `gorm:"size:50"`
	PhoneNumber string `gorm:"size:25"`
	Country     string `gorm:"size:50"`
	City        string `gorm:"size:50"`
	Ban         int64
	Mac         string `gorm:"size:17"`
}

func AddUser(id int, name string, surname string, email string, password string, phoneNumber string, country string,
	city string, mac string) string {
	database.DB.Create(&User{Id: int64(id), Name: name, Surname: surname, Mail: email,
		Password: GetMD5Hash(password), PhoneNumber: phoneNumber, Country: country, City: city, Mac: mac})
	return SelectUserName(id)
}

func SelectUserName(id int) string {
	var tbuser []User
	database.DB.Find(&tbuser)
	for _, user := range tbuser {
		if user.Id == int64(id) {
			return fmt.Sprint(user.Id) + ":" + user.Name
		}
	}
	return "nil"
}

func DeleteUser(id int64) { database.DB.Delete(&User{Id: int64(id)}) }
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
