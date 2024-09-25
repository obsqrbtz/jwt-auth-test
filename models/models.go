package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID           string `gorm:"primarykey"`
	RefreshToken string `json:"refresh_token"`
	IP           string `json:"ip"`
	Email        string `gorm:"unique"`
}

func CreateUser(ip string, email string) User {
	return User{
		ID:           uuid.NewString(),
		RefreshToken: "",
		IP:           ip,
		Email:        email,
	}
}
