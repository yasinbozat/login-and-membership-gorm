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
		tb_User     []models.User
		tb_User_Key []models.UserKey
	)

	database.DB.Find(&tb_User)
	for _, user := range tb_User {
		if user.Mail == email && user.Password == GetMD5Hash(password) {
			database.DB.Find(&tb_User_Key)
			for _, userKey := range tb_User_Key {
				if userKey.UserId == user.Id {
					if userKey.ExpiryDate.After(time.Now()) {
						RemainingTime(userKey.UserId)
						return true
					} else {
						return false
					}
				}
			}
			break
		}
	}
	return false
}

func RemainingTime(id int64) {
	var tb_userKey []models.UserKey
	database.DB.Find(&tb_userKey)
	for _, userKey := range tb_userKey {
		if userKey.UserId == id {
			if userKey.ExpiryDate.After(time.Now()) {
				fmt.Println(ParseRemainingTime(userKey.ExpiryDate))
			}
		}
	}
}

func ParseRemainingTime(expDate time.Time) string {
	var time string = expDate.Sub(CurrentTime()).String()
	var hour, minute, second string
	var x, y int
	for i := range time {

		if string(time[i]) == "h" {
			hour = time[0:i]
			y = i
		}
		if string(time[i]) == "m" {
			minute = time[y+1 : i]
			x = i
		}
		if string(time[i]) == "." {
			second = time[x+1 : i]
		}
	}
	return hour + " Hours " + minute + " Minutes " + second + " Seconds remaining..."
}

func CurrentTime() time.Time {
	var exists time.Time
	database.DB.Raw("SELECT * FROM CURRENT_TIMESTAMP;").Row().Scan(&exists)
	return exists
}
