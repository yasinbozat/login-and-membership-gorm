package main

import (
	"db_connect_app/database"
	"db_connect_app/models"
	_ "db_connect_app/models"
	"db_connect_app/utils"
	"fmt"
	"time"
)

func main() {

	//database.DB.Debug().AutoMigrate(&User{}) // Auto Migration User Table
	//database.DB.Debug().AutoMigrate(&EpinUser{})
	//AddUser(501, "Yasin", "Bozat", "admin@yasinbozat.com", "123456789", "+90 (543) 987 6543", "Turkey", "Sivas", "99:34:YB:23:BZ:58", db())
	//fmt.Print(SelectUserName(501, db()))
	utils.Login("admin@yasinbozat.com", "123456789")
	//DeleteUser(500)

}

func AddUser(id int, name string, surname string, email string, password string, phoneNumber string, country string, city string, mac string) string {

	database.DB.Create(&models.User{Id: int64(id), Name: name, Surname: surname, Mail: email, Password: utils.GetMD5Hash(password), PhoneNumber: phoneNumber, Country: country, City: city, Mac: mac})
	return SelectUserName(id)

}

func CurrentTime() time.Time {
	var exists time.Time
	database.DB.Raw("SELECT * FROM CURRENT_TIMESTAMP;").Row().Scan(&exists)

	return exists
}

func SelectUserName(id int) string {
	var tbuser []models.User
	database.DB.Find(&tbuser)
	for _, user := range tbuser {
		if user.Id == int64(id) {
			return fmt.Sprint(user.Id) + ":" + user.Name
		}
	}
	return "nil"
}

func DeleteUser(id int64) { database.DB.Delete(&models.User{Id: int64(id)}) }
