package main

import (
	"db_connect_app/models"
	_ "db_connect_app/models"
)

func main() {

	//database.DB.Debug().AutoMigrate(&models.User{},&models.Key{},&models.UserKey{}) // Auto Migration User Table
	//AddUser(501, "Yasin", "Bozat", "admin@yasinbozat.com", "123456789", "+90 (543) 987 6543", "Turkey", "Sivas", "99:34:YB:23:BZ:58", db())
	//fmt.Print(SelectUserName(501, db()))
	models.Login("admin@yasinbozat.com", "123456789")
	//DeleteUser(500)
	//utils.RemainingTime()
	//DB.Debug().AutoMigrate(&models.User{}) // Auto Migration User Table
	//DB.Debug().AutoMigrate(&models.UserKey{})
}
