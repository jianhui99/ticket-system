package ticket

import "time"

type Ticket struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	SeatNumber  string
	Price       float64
	IsAvailable bool
	DateTime    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
