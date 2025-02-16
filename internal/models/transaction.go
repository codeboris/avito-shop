package models

import "time"

type Transaction struct {
	ID        int       `db:"id"`
	FromUser  int       `db:"from_user"`
	ToUser    int       `db:"to_user"`
	Amount    int       `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
}
