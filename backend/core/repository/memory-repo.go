package repository

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/fajartd02/mygallery/backend/core/entity"
)

var ErrRecordMemoryNotFound = errors.New("record memory not found")

type MemoryRepository interface {
	Create(c *gin.Context, memory entity.Memory) (entity.Memory, error)
	Delete(c *gin.Context) error
	FindAll(c *gin.Context) ([]entity.Memory, error)
	FindByID(c *gin.Context) (entity.Memory, error)
	Update(c *gin.Context, memory entity.Memory) error
	List(c *gin.Context, memoryList entity.MemoryListRequest) ([]entity.Memory, error)
}
