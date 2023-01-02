package handler

import (
	"net/http"

	"github.com/fajartd02/mygallery/core/module"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentUc module.CommentUsecase
}

func NewCommentHandler(commentUc module.CommentUsecase) *CommentHandler {
	return &CommentHandler{
		commentUc: commentUc,
	}
}

func (hdl *CommentHandler) Create(c *gin.Context) {
	err := hdl.commentUc.CreateComment(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "comment succesfully created"})
}

func (hdl *CommentHandler) GetAll(c *gin.Context) {
	Comments, err := hdl.commentUc.GetComments(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Comments})
}
