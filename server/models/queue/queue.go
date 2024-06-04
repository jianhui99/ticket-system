package queue

import "time"

type Queue struct {
	ID         uint   `gorm:"primaryKey"`
	QueueId    string `gorm:"unique"`
	EnteredAt  time.Time
	IsPurchase bool `gorm:"default:false"`
	IpAddress  string
	BrowserId  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
