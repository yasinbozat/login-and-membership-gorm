package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Id          int64  `gorm:"primary_key autoIncrement"`
	Name        string `gorm:"size:50"`
	Surname     string `gorm:"size:50"`
	Mail        string `gorm:"size:100"`
	Password    string `gorm:"size:50"`
	PhoneNumber string `gorm:"size:25"`
	Country     string `gorm:"size:50"`
	City        string `gorm:"size:50"`
	Mac         string `gorm:"size:17"`
}

const (
	TimeFormat = "2003-08-30 15:30:00"
)

func main() {

	//db().Debug().AutoMigrate(&User{}) // Auto Migration User Table
	//AddUser(501, "Yasin", "Bozat", "admin@yasinbozat.com", "123456789", "+90 (531) 833 2425", "Turkey", "Sivas", "99:34:YB:23:BZ:58", db())
	//fmt.Print(SelectUserName(501, db()))
	fmt.Println(Login("admin@yasinbozat.com", "123456789"))

}

func AddUser(id int, name string, surname string, email string, password string, phoneNumber string, country string, city string, mac string, db *gorm.DB) string {

	db.Create(&User{Id: int64(id), Name: name, Surname: surname, Mail: email, Password: GetMD5Hash(password), PhoneNumber: phoneNumber, Country: country, City: city, Mac: mac})
	return SelectUserName(id, db)

}

func Login(email, password string) bool {
	var tbuser []User
	db().Find(&tbuser)
	for _, user := range tbuser {
		if user.Mail == email && user.Password == GetMD5Hash(password) {
			return true
		}
	}
	return false
}

func CurrentTime() time.Time {
	var exists time.Time
	db().Raw("SELECT * FROM CURRENT_TIMESTAMP;").Row().Scan(&exists)

	return exists
}

func SelectUserName(id int, db *gorm.DB) string {
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

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
