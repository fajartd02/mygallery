package memoryrepository

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fajartd02/mygallery/core/entity"
	repository_intf "github.com/fajartd02/mygallery/core/repository"
)

type repositoryMemory struct {
}

func New() repository_intf.MemoryRepository {
	return &repositoryMemory{}
}

func (r *repositoryMemory) Create(c *gin.Context, memory entity.Memory) (entity.Memory, error) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return entity.Memory{}, errors.New("failed to parse db to gorm")
	}

	if err := db.Create(&memory).Error; err != nil {
		return entity.Memory{}, errors.New("failed to create memory")
	}

	return memory, nil
}

func (r *repositoryMemory) FindAll(c *gin.Context) ([]entity.Memory, error) {
	var memories []entity.Memory

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return nil, errors.New("failed to parse db to gorm")
	}

	userIdHeader := c.Request.Header["User-Id"]
	if err := db.Model(&entity.Memory{}).Where("user_id = ?", userIdHeader).Find(&memories).Error; err != nil {
		return nil, repository_intf.ErrRecordUserNotFound
	}

	return memories, nil
}

func (r *repositoryMemory) List(c *gin.Context, memoryList entity.MemoryListRequest) ([]entity.Memory, error) {
	var memories []entity.Memory

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return nil, errors.New("failed to parse db to gorm")
	}
	userIdHeader := c.Request.Header["User-Id"]

	filterQuery := "user_id = " + userIdHeader[0] + " AND LOWER(" + memoryList.Filter.By + ") LIKE LOWER(?)"
	sortQuery := memoryList.Sort.By + " " + memoryList.Sort.Order

	if err := db.Model(&entity.Memory{}).Where(filterQuery, "%"+memoryList.Filter.Keyword+"%").Order(sortQuery).Find(&memories).Error; err != nil {
		return nil, repository_intf.ErrRecordUserNotFound
	}
	return memories, nil
}

func (r *repositoryMemory) Update(c *gin.Context, memory entity.Memory) error {
	db, ok := c.MustGet("db").(*gorm.DB)

	if !ok {
		return errors.New("failed to parse db to gorm")
	}
	var oldMemory entity.Memory
	if err := db.First(&oldMemory, c.Param("memoryId")).Error; err != nil {
		return errors.New("old memory didn't exists!")
	}

	if err := db.Where("id = ?", c.Param("memoryId")).Updates(&memory).Error; err != nil {
		return errors.New("failed to update memory")
	}

	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	deleteDir := mydir + "/images/" + oldMemory.ImageUrl
	e := os.Remove(deleteDir)
	if e != nil {
		log.Fatal(e)

	}

	return nil
}

func (r *repositoryMemory) Delete(c *gin.Context) error {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return errors.New("failed to parse db to gorm")
	}

	var memory entity.Memory
	if err := db.Where("id = ?", c.Param("memoryId")).First(&memory).Error; err != nil {
		return errors.New("failed to delete user")
	}

	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	deleteDir := mydir + "/images/" + memory.ImageUrl
	fmt.Println(deleteDir)
	e := os.Remove(deleteDir)
	if e != nil {
		log.Fatal(e)

	}

	if err := db.Delete(&memory).Error; err != nil {
		return errors.New("failed to delete memory")
	}

	return nil
}
