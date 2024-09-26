package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID           string `gorm:"primarykey:not null"`
	RefreshToken string `json:"refresh_token"`
	IP           string `json:"ip"`
	Email        string `gorm:"unique:not null"`
}

func CreateUser(ip string, email string) User {
	return User{
		ID:           uuid.NewString(),
		RefreshToken: "",
		IP:           ip,
		Email:        email,
	}
}
