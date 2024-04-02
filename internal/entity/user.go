package entity

import (
	"time"
)

type User struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
}

type CreateUser struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UpdateUser struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
