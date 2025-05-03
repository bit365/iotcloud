package models

import "time"

type User struct {
	ID        uint
	UserName  string
	Password  string
	Email     string
	Phone     string
	CreatedAt time.Time
}
