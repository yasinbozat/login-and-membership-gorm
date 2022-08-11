package models

import (
	"db_connect_app/database"
	"fmt"
	"time"
)

const (
	DDMMYYYYhhmmss = "2006-01-02 15:04:05"
	DDhhmmss       = "02 15:04:05"
)

func RemainingTime(id int64) {
	var tb_userKey []UserKey
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
