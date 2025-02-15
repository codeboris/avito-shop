package models

import "time"

type Transaction struct {
	ID        uint
	FromUser  uint
	ToUser    uint
	Amount    int
	CreatedAt time.Time
}
