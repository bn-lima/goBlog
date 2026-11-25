package model

import "time"

type User struct {
	ID        uint
	Username  string
	CreatedAt time.Time
	Email     string
	Password  string
}
