package user

import (
	"time"
)

type User struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Email       string
	PhoneNumber string
	Password    string
}
