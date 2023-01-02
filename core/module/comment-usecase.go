package module

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/fajartd02/mygallery/core/entity"
	"github.com/fajartd02/mygallery/core/repository"
)

type CommentUsecase interface {
	CreateComment(c *gin.Context) error
	GetComments(c *gin.Context) ([]entity.Comment, error)
}

type commentUsecase struct {
	commentRepo repository.CommentRepository
	userRepo    repository.UserRepository
}

var ErrCommentNotFound = errors.New("comment error: ")

func NewCommentUsecase(repo repository.CommentRepository, userrepo repository.UserRepository) CommentUsecase {
	return &commentUsecase{repo, userrepo}
}

func (em *commentUsecase) CreateComment(c *gin.Context) error {
	userIdHeader := c.Request.Header["User-Id"]
	if len(userIdHeader) == 0 {
		return errors.New("failed to parse userID")
	}

	userID, err := strconv.ParseUint(userIdHeader[0], 10, 32)
	if err != nil {
		return errors.New("failed to parse userID")
	}

	user, err := em.userRepo.FindSingle(c, uint(userID))
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}

	err = em.commentRepo.Create(c, uint(userID), user.FullName)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCommentNotFound, err)
	}
	return nil
}

func (em *commentUsecase) GetComments(c *gin.Context) ([]entity.Comment, error) {
	data, err := em.commentRepo.FindAll(c)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCommentNotFound, err)
	}

	return data, nil
}
