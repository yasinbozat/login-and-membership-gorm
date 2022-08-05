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

	var dsn = os.Getenv("CONNECTION_STRING")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.Debug().AutoMigrate(&User{})
	db.Create(&User{Id: 1, Name: "Yasin"})

	var tbuser []User
	db.Find(&tbuser)
	for _, user := range tbuser {
		fmt.Printf("ID:%d\nNickname:%s\n", user.Id, user.Name)
		fmt.Printf("--------------------------------------------------------------\n")
	}
}
