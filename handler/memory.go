package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/fajartd02/mygallery/core/entity"
	"github.com/fajartd02/mygallery/core/module"
)

type MemoryHandler struct {
	memoryUc module.MemoryUseCase
}

func NewMemoryHandler(memoryUc module.MemoryUseCase) *MemoryHandler {
	return &MemoryHandler{
		memoryUc: memoryUc,
	}
}

func (hdl *MemoryHandler) CreateMemory(c *gin.Context) {
	var input entity.MemoryInput
	if err := c.Bind(&input); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	userIdHeader := c.Request.Header["User-Id"]
	if len(userIdHeader) == 0 {
		c.AbortWithStatusJSON(500, gin.H{"message": "userID not found in the header!"})
		return
	}

	userID, err := strconv.ParseUint(userIdHeader[0], 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	file, _ := c.FormFile("file")
	var newNameFile = uuid.NewString() + "-" + file.Filename
	var dst = "images/" + newNameFile
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	// Create Memory
	memory := entity.Memory{
		UserID:      uint(userID),
		ImageUrl:    newNameFile,
		Description: input.Description,
		Tag:         input.Tag,
	}

	err = hdl.memoryUc.CreateMemory(c, memory)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sucessfully created memory!",
		"status":  http.StatusOK,
		"data":    memory,
	})
}

func (hdl *MemoryHandler) GetAll(c *gin.Context) {
	memories, err := hdl.memoryUc.FindAll(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sucessfully fetch all memory!",
		"status":  http.StatusOK,
		"data":    memories,
	})
}

func (hdl *MemoryHandler) ListMemory(c *gin.Context) {
	var input entity.MemoryListRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
	}

	memories, err := hdl.memoryUc.List(c, input)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sucessfully fetch all memory!",
		"status":  http.StatusOK,
		"data":    memories,
	})
}

func (hdl *MemoryHandler) DeleteMemory(c *gin.Context) {
	err := hdl.memoryUc.DeleteMemory(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sucessfully deleted memory!",
		"status":  http.StatusOK,
	})
}

func (hdl *MemoryHandler) UpdateMemory(c *gin.Context) {
	var input entity.MemoryInput
	if err := c.Bind(&input); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	file, _ := c.FormFile("file")
	var newNameFile = uuid.NewString() + "-" + file.Filename
	var dst = "images/" + newNameFile

	// Update User
	memory := entity.Memory{
		ImageUrl:    newNameFile,
		Description: input.Description,
		Tag:         input.Tag,
	}

	err := hdl.memoryUc.UpdateMemory(c, memory)

	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sucessfully updated memory!",
		"status":  http.StatusOK,
	})
}
