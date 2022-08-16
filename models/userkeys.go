package models

import (
	"db_connect_app/database"
	"fmt"
	"time"
)

type UserKey struct {
	Id            int64 `gorm:"primary_key;auto_increment"`
	UserId        int64
	KeyId         int64
	UsingDate     time.Time
	ExpiryDate    time.Time
	RemainingTime RemTime `gorm:"-"`
}

type RemTime struct {
	Day, Hour, Minute, Second int
}

func UseKey(email, key string) {

	var (
		keys = Key{Key: key}
		user = User{Mail: email}
		userkeys = UserKey{}
	)
	if results := database.DB.Where(&user, "mail").First(&user); results.Error == nil { 
		if results = database.DB.Where(&keys, "key").First(&keys); results.Error == nil { 
			if keys.Active == 1 && keys.Ban == 0 {
				if user.Ban == 0 {
					if results = database.DB.Where(&userkeys, "us").First(&userkeys); results.Error == nil { 

					}
				} else {
					fmt.Println("You cannot use a key on a blocked account.")
					return
				}
			} else {
				fmt.Println("Invalid key!")
				return
			}
		} else {
			fmt.Println("Invalid key!")
			return
		}
	} else {
		fmt.Println("Invalid username!")
		return
	}

}
