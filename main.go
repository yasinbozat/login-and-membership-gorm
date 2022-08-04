package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Id   int64  `gorm:"primary_key"`
	Name string `gorm:"size:255"`
}

var dsn = "host=localhost user=postgres password=1234 dbname=db_user port=5432 sslmode=disable TimeZone=Asia/Istanbul"

func main() {

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
