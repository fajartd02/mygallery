package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	ID           uuid.UUID
	Filename     string
	OriginalName string
	Mimetype     string
	Url          string
	Size         int64
	Extension    string
}

type FileInput struct {
	Filename     string `json:"filename"`
	OriginalName string `json:"original_name"`
	Mimetype     string `json:"mimetype"`
	Url          string `json:"url"`
	Size         int64  `json:"size"`
	Extension    string `json:"extension"`
}
