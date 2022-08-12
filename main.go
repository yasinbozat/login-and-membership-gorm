package main

import (
	"db_connect_app/database"
	"db_connect_app/models"
	_ "db_connect_app/models"
	"db_connect_app/utils"
	"fmt"
)

func main() {

	database.DB.Debug().AutoMigrate(&models.User{}, &models.Key{}, &models.UserKey{}) // Auto Migration User Table
	//models.Login("admin@yasinbozat.com", "123456789")
	fmt.Println(utils.GetHWID())
}
