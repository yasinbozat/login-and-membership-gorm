package models

import (
	"time"
)

type UserKey struct {
	Id            int64 `gorm:"primary_key autoIncrement"`
	UserId        int64
	KeyId         int64
	UsingDate     time.Time
	ExpiryDate    time.Time
	RemainingTime RemTime `gorm:"-"`
}

type RemTime struct {
	Day, Hour, Minute, Second int
}
