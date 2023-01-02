package handler

import (
	"errors"
	"mime"
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/fajartd02/mygallery/core/entity"
)

func parseFileInput(c *gin.Context) (entity.File, multipart.File, error) {
	file, err := c.FormFile("fileUpload")
	if err != nil {
		return entity.File{}, nil, err
	}

	// Get Buffer from file
	buffer, err := file.Open()

	if err != nil {
		return entity.File{}, nil, err
	}

	var maxSize = 10 << (10 * 2)

	if file.Size > int64(maxSize) {
		return entity.File{}, nil, errors.New("file too large")
	}

	ext := filepath.Ext(file.Filename)
	return entity.File{
		Filename:     file.Filename,
		OriginalName: file.Filename,
		Extension:    ext,
		Mimetype:     mime.TypeByExtension(ext),
		Size:         file.Size,
	}, buffer, nil
}
