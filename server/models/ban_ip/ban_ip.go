package banip

import "time"

type BanIp struct {
	Id        int `json:"id"`
	IpAddress string
	CreatedAt time.Time
}
