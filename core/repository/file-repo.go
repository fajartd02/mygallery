package repository

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/fajartd02/mygallery/core/entity"
)

var ErrRecordFileNotFound = errors.New("record not found")

type FileRepository interface {
	FindAll(c *gin.Context) ([]entity.File, error)
	FindSingle(c *gin.Context, filename string) (entity.File, error)
	Create(c *gin.Context, file *entity.File) (entity.File, error)
	Delete(c *gin.Context, filename string) error
}
