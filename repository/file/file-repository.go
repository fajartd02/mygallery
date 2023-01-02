package filerepository

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	repository_intf "github.com/fajartd02/mygallery/core/repository"

	"github.com/fajartd02/mygallery/core/entity"
)

var ErrDBNotFound = errors.New("repo: db not found")

type repository struct {
}

func New() repository_intf.FileRepository {
	return &repository{}
}

func (r *repository) FindAll(c *gin.Context) ([]entity.File, error) {
	var files []entity.File

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return files, errors.New("failed to parse db to gorm")
	}

	err := db.Find(&files).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return files, nil
}

func (r *repository) FindSingle(c *gin.Context, filename string) (entity.File, error) {
	file := entity.File{}
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return file, errors.New("failed to parse db to gorm")
	}

	err := db.Where("filename = ?", filename).First(&file).Error
	if err != nil {
		return entity.File{}, err
	}

	return file, nil
}

func (r *repository) Create(c *gin.Context, file *entity.File) (entity.File, error) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return entity.File{}, errors.New("failed to parse db to gorm")
	}

	if err := db.Create(&file).Error; err != nil {
		return entity.File{}, err
	}

	return *file, nil
}

func (r *repository) Delete(c *gin.Context, filename string) error {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return errors.New("failed to parse db to gorm")
	}

	var file entity.File
	if err := db.Where("filename = ?", filename).First(&file).Error; err != nil {
		return err
	}

	if err := db.Delete(&file).Error; err != nil {
		return err
	}

	return nil
}
