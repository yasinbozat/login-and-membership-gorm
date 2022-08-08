package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Id   int64  `gorm:"primary_key"`
	Name string `gorm:"size:255"`
}

const (
	TimeFormat = "2003-08-30 15:30:00"
)

func main() {
	db().Debug().AutoMigrate(&User{}) // Auto Migration User Table

	fmt.Println(SelectUser(2, db()))              // Select user
	fmt.Println(CurrentTime().Format(TimeFormat)) // Get db time
	AddUser(8, "Mahmut Tuncer", db())             // Add user
}

func AddUser(id int, name string, db *gorm.DB) string {

	db.Create(&User{Id: int64(id), Name: string(name)})
	return SelectUser(id, db)

}

func CurrentTime() time.Time {
	var exists time.Time
	db().Raw("SELECT * FROM CURRENT_TIMESTAMP;").Row().Scan(&exists)

	return exists
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

func db() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("CONNECTION_STRING")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db

}
