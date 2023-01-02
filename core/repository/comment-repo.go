package repository

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/fajartd02/mygallery/core/entity"
)

var ErrRecordCommentNotFound = errors.New("record report not found")

type CommentRepository interface {
	FindAll(c *gin.Context) ([]entity.Comment, error)
	Create(c *gin.Context, id uint, fullname string) error
}
