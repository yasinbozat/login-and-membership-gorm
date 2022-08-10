package utils

import (
	"crypto/md5"
	"db_connect_app/database"
	"db_connect_app/models"
	"encoding/hex"
	"fmt"
	"time"
)

const (
	DDMMYYYYhhmmss = "2006-01-02 15:04:05"
	DDhhmmss       = "02 15:04:05"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func Login(email, password string) bool {
	var (
		tbuser    []models.User
		customers []models.EpinUser
		UserId    int64
		name      string
	)

	database.DB.Find(&tbuser)
	for _, user := range tbuser {
		if user.Mail == email && user.Password == GetMD5Hash(password) {
			UserId = user.Id
			name = user.Name + " " + user.Surname
			break
		}
	}

	database.DB.Find(&customers)
	for _, epinUser := range customers {
		epinUser.RemainingTime.Day = epinUser.ExpiryDate.Day() - time.Now().Day()
		epinUser.RemainingTime.Hour = epinUser.ExpiryDate.Hour() - time.Now().Hour()
		epinUser.RemainingTime.Minute = epinUser.ExpiryDate.Minute() - time.Now().Minute()
		epinUser.RemainingTime.Second = epinUser.ExpiryDate.Minute() - time.Now().Second()
		if epinUser.UserId == UserId {
			if epinUser.ExpiryDate.After(time.Now()) {
				fmt.Printf("Hello %v,\nYour expiry date: %v\n%v Days %v Hours %v Minutes %v Seconds", name,
																			epinUser.ExpiryDate.Format(DDMMYYYYhhmmss), 
																			epinUser.RemainingTime.Day, 
																			epinUser.RemainingTime.Hour, 
																			epinUser.RemainingTime.Minute, 
																			epinUser.RemainingTime.Second)
				return true
			} else {
				return false
			}
		}
	}

	return false
}
