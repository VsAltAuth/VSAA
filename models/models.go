package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	UID          string
	Email        string
	HashedPass   string
	Salt         string
	Playername   string
	Entitlements string
}

type Session struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	UID        string
	Sessionkey string
	Gamever    string
}
