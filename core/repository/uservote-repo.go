package repository

import (
	"errors"

	"github.com/fajartd02/mygallery/core/entity"
	"github.com/gin-gonic/gin"
)

var ErrRecordUserVoteNotFound = errors.New("record vote not found")

type UserVoteRepository interface {
	ChooseVotes(c *gin.Context) error
	GetVotesInReport(c *gin.Context) (entity.Votes, error)
}
