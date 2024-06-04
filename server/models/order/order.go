package order

import "time"

type Order struct {
	ID        uint `gorm:"primaryKey"`
	Email     string
	TicketId  uint
	Quantity  uint
	Total     float64
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
