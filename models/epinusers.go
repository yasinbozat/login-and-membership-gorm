package models

import (
	"time"
)

type EpinUser struct {
	Id            int64 `gorm:"primary_key autoIncrement"`
	UserId        int64
	EpinId        int64
	UsingDate     time.Time
	ExpiryDate    time.Time
	RemainingTime RemTime `gorm:"-"`
}

type RemTime struct {
	Day, Hour, Minute, Second int
}
