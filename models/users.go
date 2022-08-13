package models

import (
	"db_connect_app/database"
	"db_connect_app/utils"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
)

type User struct {
	Id          int64  `gorm:"primaryKey;autoIncrement:true"`
	Name        string `gorm:"size:50"`
	Surname     string `gorm:"size:50"`
	Mail        string `gorm:"size:100;not_null;unique"`
	Password    string `gorm:"size:50;not_null"`
	PhoneNumber string `gorm:"size:25"`
	Country     string `gorm:"size:50"`
	City        string `gorm:"size:50"`
	Ban         byte
	HWID        string `gorm:"size:33;not_null"`
}

func AddUser(id int, name string, surname string, email string, password string, phoneNumber string, country string,
	city string, hwid string) string {
	database.DB.Create(&User{Id: int64(id), Name: name, Surname: surname, Mail: email,
		Password: utils.GetMD5Hash(password), PhoneNumber: phoneNumber, Country: country, City: city, HWID: hwid})
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
	if results := database.DB.Where(&user, "mail", "password").First(&user); results.Error == nil { //Check mail and password
		if user.Ban == 0 {
			if user.HWID == utils.GetHWID() {
				var userKey = UserKey{UserId: user.Id}
				//Find remaining time from user key using user id
				if results = database.DB.Where(&userKey, "user_id").First(&userKey); results.Error == nil {
					if userKey.ExpiryDate.After(CurrentTime()) { //Check expiry date > current time
						RemainingTime(userKey.UserId)
						return true
					} else {
						fmt.Println("Your time has expired! Please activate a key.") //If user time has expired
						return false
					}
				} else {
					fmt.Println("Please activate a key.") //If user never used key
					return false
				}
			} else {
				fmt.Println("Logged in from an unknown computer. Please login from the registered computer.") //If user has wrong hwid
				return false
			}
		} else {
			fmt.Println("Your account has been banned!") //If user has been banned
			return false
		}
	} else if user.Id == 0 {
		fmt.Println("Invalid username or password.")
		return false
	}
	return false

}

func CreateUser(name, surname, mail, password, phoneNumber, country, city string) {

	user := User{Name: name, Surname: surname, Mail: mail, Password: utils.GetMD5Hash(password), PhoneNumber: phoneNumber,
		Country: country, City: city, Ban: 0, HWID: utils.GetHWID()}

	if results := database.DB.Create(&user); results.Error != nil {
		if pgError := results.Error.(*pgconn.PgError); errors.Is(results.Error, pgError) {
			switch pgError.Code {
			case "23505":
				fmt.Println("")
			}
		}
	}

}
