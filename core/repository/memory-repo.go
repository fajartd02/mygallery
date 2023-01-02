package repository

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/fajartd02/mygallery/core/entity"
)

var ErrRecordMemoryNotFound = errors.New("record memory not found")

type MemoryRepository interface {
	Create(c *gin.Context, memory entity.Memory) (entity.Memory, error)
	Delete(c *gin.Context) error
	FindAll(c *gin.Context) ([]entity.Memory, error)
	Update(c *gin.Context, memory entity.Memory) error
}
