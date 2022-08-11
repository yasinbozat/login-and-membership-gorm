package main

import (
	"db_connect_app/database"
	"db_connect_app/models"
	_ "db_connect_app/models"
	"db_connect_app/utils"
)

func main() {

	database.DB.Debug().AutoMigrate(&models.User{}) // Auto Migration User Table
	database.DB.Debug().AutoMigrate(&models.UserKey{})
	database.DB.Debug().AutoMigrate(&models.Key{})
	//AddUser(501, "Yasin", "Bozat", "admin@yasinbozat.com", "123456789", "+90 (543) 987 6543", "Turkey", "Sivas", "99:34:YB:23:BZ:58", db())
	//fmt.Print(SelectUserName(501, db()))
	utils.Login("admin@yasinbozat.com", "123456789")
	//DeleteUser(500)
	//utils.RemainingTime()
	//DB.Debug().AutoMigrate(&models.User{}) // Auto Migration User Table
	//DB.Debug().AutoMigrate(&models.UserKey{})
}
