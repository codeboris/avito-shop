package models

import "time"

type Purchase struct {
	ID        uint
	UserID    uint
	MerchID   uint
	Quantity  int
	CreatedAt time.Time
}
