package user

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Email     string
	Password  string
	IpAddress string
	BrowserId string
	CreatedAt time.Time
	UpdatedAt time.Time
}
