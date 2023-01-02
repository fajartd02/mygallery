package repository

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/fajartd02/mygallery/core/entity"
)

var ErrRecordUserNotFound = errors.New("record not found")

type UserRepository interface {
	FindAll(c *gin.Context) ([]entity.User, error)
	FindSingle(c *gin.Context, id uint) (entity.User, error)
	FindFullNames(c *gin.Context, userIDs []uint) ([]string, error)
	FindSingleByEmail(c *gin.Context, email string) (entity.User, error)
	Create(c *gin.Context, user entity.User) error
	Update(c *gin.Context, user entity.User) error
	Delete(c *gin.Context) error
}
