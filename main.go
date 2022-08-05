package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Id   int64  `gorm:"primary_key"`
	Name string `gorm:"size:255"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("CONNECTION_STRING")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	//db.Debug().AutoMigrate(&User{})

	fmt.Println(AddUser(5, "Hakan", db))

}

func AddUser(id int, name string, db *gorm.DB) string {

	db.Create(&User{Id: int64(id), Name: string(name)})

	return SelectUser(id, db)
}

func SelectUser(id int, db *gorm.DB) string {
	var tbuser []User
	db.Find(&tbuser)
	for _, user := range tbuser {
		if user.Id == int64(id) {
			return fmt.Sprint(user.Id) + ":" + user.Name
		}
	}
	return "nil"
}
