package module

import (
	"errors"
	"fmt"

	"github.com/fajartd02/mygallery/core/entity"
	"github.com/fajartd02/mygallery/core/repository"
	"github.com/gin-gonic/gin"
)

type UserVoteUsecase interface {
	VotingReport(c *gin.Context) error
	GetVotesInReport(c *gin.Context) (entity.Votes, error)
}

type userVoteUsecase struct {
	userVoteRepo repository.UserVoteRepository
}

var ErrUserVoteNotFound = errors.New("uservote error: ")

func NewUserVoteUsecase(repo repository.UserVoteRepository) UserVoteUsecase {
	return &userVoteUsecase{repo}
}

func (em *userVoteUsecase) VotingReport(c *gin.Context) error {
	err := em.userVoteRepo.ChooseVotes(c)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUserVoteNotFound, err)
	}
	return nil
}

func (em *userVoteUsecase) GetVotesInReport(c *gin.Context) (entity.Votes, error) {
	data, err := em.userVoteRepo.GetVotesInReport(c)
	if err != nil {
		return entity.Votes{}, fmt.Errorf("%w: %v", ErrUserVoteNotFound, err)
	}
	return data, nil
}
