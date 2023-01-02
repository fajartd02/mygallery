package handler

import (
	"net/http"

	"github.com/fajartd02/mygallery/core/module"
	"github.com/gin-gonic/gin"
)

type UserVoteHandler struct {
	userVoteUc module.UserVoteUsecase
}

func NewUserVoteHandler(userVoteUc module.UserVoteUsecase) *UserVoteHandler {
	return &UserVoteHandler{
		userVoteUc: userVoteUc,
	}
}

func (hdl *UserVoteHandler) VotingReport(c *gin.Context) {
	err := hdl.userVoteUc.VotingReport(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "voting succesfully updated or created"})
}

func (hdl *UserVoteHandler) GetVotesInReport(c *gin.Context) {
	Vote, err := hdl.userVoteUc.GetVotesInReport(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Vote})
}
