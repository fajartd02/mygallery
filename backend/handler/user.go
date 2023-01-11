package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fajartd02/mygallery/backend/core/entity"
	"github.com/fajartd02/mygallery/backend/core/module"
)

type UserHandler struct {
	userUc module.UserUsecase
}

func NewUserHandler(userUc module.UserUsecase) *UserHandler {
	return &UserHandler{
		userUc: userUc,
	}
}

func (hdl *UserHandler) GetAll(c *gin.Context) {
	Users, err := hdl.userUc.GetUsers(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Users})
}

func (hdl *UserHandler) GetSingle(c *gin.Context) {
	User, err := hdl.userUc.GetUser(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": User})
}

func (hdl *UserHandler) Login(c *gin.Context) {
	var input entity.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	User, err := hdl.userUc.Login(c, input)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": User})
}

func (hdl *UserHandler) Register(c *gin.Context) {
	var input entity.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	// Register User
	User := entity.User{
		FullName: input.FullName,
		Email:    input.Email,
		Password: input.Password,
	}

	err := hdl.userUc.CreateUser(c, User)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "user successfully created"})
}

func (hdl *UserHandler) Update(c *gin.Context) {
	var input entity.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	// Update User
	User := entity.User{
		FullName: input.FullName,
		Email:    input.Email,
		Password: input.Password,
	}

	err := hdl.userUc.UpdateUser(c, User)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "user successfully updated"})
}
func (hdl *UserHandler) Delete(c *gin.Context) {
	err := hdl.userUc.DeleteUser(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "user successfully deleted"})
}
