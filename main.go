package main

import (
	"db_connect_app/database"
	"db_connect_app/models"
	_ "db_connect_app/models"
)

func main() {

	database.DB.Debug().AutoMigrate(&models.User{}, &models.Key{}, &models.UserKey{}) // Auto Migration User Table
	//models.Login("admin@yasinbozat.com", "123456789")
	//fmt.Println(utils.GetHWID())
	models.CreateUser("Yasin2", "Bozat2", "yasinbozatr2@gmail.com", "123456789", "5318332425", "Turkey", "Istanbul")
}
