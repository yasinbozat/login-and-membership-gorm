package models

import (
	"time"
)

type Key struct {
	Id         int64  `gorm:"primary_key;auto_increment"`
	Key        string `gorm:"size:32;unique"`
	Type       byte
	Active     byte
	Day        int
	Ban        byte
	CreateDate time.Time
}
