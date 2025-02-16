package models

import "time"

type User struct {
	ID        int       `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Coins     int       `db:"coins"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
