package main

import (
	"db_connect_app/models"
	_ "db_connect_app/models"
)

func main() {

	//database.DB.Debug().AutoMigrate(&models.User{}, &models.Key{}, &models.UserKey{}) // Auto Migration User Table
	//models.Login("yasinbozatr@gmail.com", "123456789")
	models.UseKey("yasinbozatr@gmail.com", "17877-AIR76-MU85N-3Y84U-P78XX")
	//fmt.Println(utils.GetHWID())
	//models.CreateUser("Yasin", "Bozat", "yasinbozatr@gmail.com", "123456789", "5321234567", "Turkey", "Istanbul")
	//fmt.Println(models.CurrentTime())
}
