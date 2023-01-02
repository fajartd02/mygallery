package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fajartd02/mygallery/core/entity"
	"github.com/fajartd02/mygallery/core/module"
)

type FileHandler struct {
	fileUc module.FileUsecase
}

func NewFileHandler(fileUc module.FileUsecase) *FileHandler {
	return &FileHandler{
		fileUc,
	}
}
func (hdl *FileHandler) UploadFile(c *gin.Context) {
	var file entity.File

	file, buffer, err := parseFileInput(c)
	if err != nil {
		log.Println("[Upload file] : ", err)
		c.AbortWithStatusJSON(400, err)
		return
	}
	defer buffer.Close()

	info, err := hdl.fileUc.UploadFileToStorage(c, file, buffer)
	if err != nil {
		log.Println("[Upload file] : ", err)
		c.AbortWithStatusJSON(500, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": info})
}
