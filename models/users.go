package models

import (
	"db_connect_app/database"
	"db_connect_app/utils"
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
	Ban         byte
	Mac         string `gorm:"size:17"`
}

func AddUser(id int, name string, surname string, email string, password string, phoneNumber string, country string,
	city string, mac string) string {
	database.DB.Create(&User{Id: int64(id), Name: name, Surname: surname, Mail: email,
		Password: utils.GetMD5Hash(password), PhoneNumber: phoneNumber, Country: country, City: city, Mac: mac})
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

func Login(email, password string) bool {
	var user = User{Mail: email, Password: utils.GetMD5Hash(password)}
	database.DB.Where(&user, "mail", "password").First(&user) //Check mail and password
	if user.Id != 0 && user.Ban == 0 {
		var userKey = UserKey{UserId: user.Id}
		database.DB.Where(&userKey, "user_id").First(&userKey) //Find remaining time from user key using user id
		if userKey.Id != 0 {
			if userKey.ExpiryDate.After(CurrentTime()) { //Check expiry date > current time
				RemainingTime(userKey.UserId)
				return true
			} else {
				fmt.Println("Your time has expired! Please activate a key.")
				return false
			}
		} else {
			fmt.Println("Please activate a key.")
			return false
		}
	} else if user.Id == 0 {
		fmt.Println("Username or password incorrect.")
		return false
	}
	return false
}
