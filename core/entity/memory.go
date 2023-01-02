package entity

import (
	"gorm.io/gorm"
)

type Memory struct {
	gorm.Model
	UserID      uint
	Description string `json:"description"`
	Tag         string `json:"tag"`
	ImageUrl    string `json:"imageUrl"`
}

type MemoryInput struct {
	Description string `form:"description"`
	Tag         string `form:"tag"`
	ImageUrl    string `form:"imageUrl"`
}
