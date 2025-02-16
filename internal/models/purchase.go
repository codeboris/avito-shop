package models

import "time"

type Purchase struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	MerchID   int       `db:"merch_id"`
	Quantity  int       `db:"quantity"`
	CreatedAt time.Time `db:"created_at"`
}
