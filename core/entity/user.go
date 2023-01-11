package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Memories []Memory `json:"Memories"`
	FullName string   `json:"fullName"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
}

type UserInput struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	UserID      uint   `json:"userID"`
	TokenString string `json:"token"`
}
