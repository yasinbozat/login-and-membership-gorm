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

func UseKey(email, key string) bool {

	var (
		keys     = Key{Key: key}
		user     = User{Mail: email}
		userkeys = UserKey{}
	)

	if results := database.DB.Where(&user, "mail").First(&user); results.Error == nil {
		userkeys.UserId = user.Id
		if user.Ban == 0 {
			if results = database.DB.Where(&keys, "key").First(&keys); results.Error == nil {
				if keys.Active == 1 && keys.Ban == 0 {
					if results = database.DB.Where(&userkeys, "user_id").First(&userkeys); results.Error == nil {
						database.DB.Where("user_id = ?", user.Id).Delete(&userkeys)
						AddDay(userkeys, user, keys)
						return true
					} else {
						AddDay(userkeys, user, keys)
						return true
					}
				} else {
					fmt.Println("Invalid key!")
					return false
				}
			} else {
				fmt.Println("Invalid key.")
				return false
			}
		} else {
			fmt.Println("You cannot use a key on a blocked account.")
			return false
		}
	} else {
		fmt.Println("Invalid username!")
		return false
	}

	return false
}

func AddDay(userkeys UserKey, user User, keys Key) {
	userkeys.ExpiryDate = CalculateExpiryDate(userkeys, keys)
	userkeys = UserKey{
		UserId:     user.Id,
		KeyId:      keys.Id,
		UsingDate:  CurrentTime(),
		ExpiryDate: userkeys.ExpiryDate}
	database.DB.Create(&userkeys)
	err := database.DB.Model(&keys).Update("active", 0).Error
	if err != nil {
		fmt.Println(err)
	}
}
