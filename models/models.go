package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UID          string
	Email        string
	HashedPass   string
	Playername   string
	Entitlements string
}

type Session struct {
	gorm.Model
	UID        string
	Sessionkey string
	Gamever    string
}
