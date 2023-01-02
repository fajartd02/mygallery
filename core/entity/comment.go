package entity

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	FullName    string
	UserID      uint
	ReportID    uint
	Description string `json:"description"`
}

type CommentResp struct {
	gorm.Model
	UserID      uint
	ReportID    uint
	Description string `json:"description"`
}

type CommentInput struct {
	Description string `json:"description"`
}
