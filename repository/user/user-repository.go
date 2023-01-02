package userrepository

import (
	"errors"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"github.com/fajartd02/mygallery/core/entity"
	repository_intf "github.com/fajartd02/mygallery/core/repository"
)

type repositoryUser struct {
}

func New() repository_intf.UserRepository {
	return &repositoryUser{}
}

func (r *repositoryUser) FindFullNames(c *gin.Context, userIDs []uint) ([]string, error) {
	var users []entity.User

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return nil, errors.New("failed to parse db to gorm")
	}

	err := db.Find(&users, userIDs).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrRecordUserNotFound
		}
		return nil, err
	}

	var fullNames []string
	for _, el := range users {
		fullNames = append(fullNames, el.FullName)
	}

	return fullNames, nil
}

func (r *repositoryUser) FindAll(c *gin.Context) ([]entity.User, error) {
	var users []entity.User

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return nil, errors.New("failed to parse db to gorm")
	}

	err := db.Find(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrRecordUserNotFound
		}
		return nil, err
	}

	return users, nil
}
func (r *repositoryUser) FindByID(c *gin.Context) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repositoryUser) FindSingle(c *gin.Context, id uint) (entity.User, error) {
	user := entity.User{}

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return entity.User{}, errors.New("failed to parse db to gorm")
	}

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return entity.User{}, repository_intf.ErrRecordUserNotFound
	}

	return user, nil
}

func (r *repositoryUser) FindSingleByEmail(c *gin.Context, email string) (entity.User, error) {
	user := entity.User{}

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return entity.User{}, errors.New("failed to parse db to gorm")
	}

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return entity.User{}, repository_intf.ErrRecordUserNotFound
	}

	return user, nil
}

func (r *repositoryUser) Create(c *gin.Context, user entity.User) error {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return errors.New("failed to parse db to gorm")
	}

	if err := db.Create(&user).Error; err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (r *repositoryUser) Update(c *gin.Context, user entity.User) error {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return errors.New("failed to parse db to gorm")
	}

	if err := db.Where("id = ?", c.Param("id")).Updates(&user).Error; err != nil {
		return errors.New("failed to update user")
	}

	return nil
}

func (r *repositoryUser) Delete(c *gin.Context) error {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return errors.New("failed to parse db to gorm")
	}

	var user entity.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		return errors.New("failed to delete user")
	}

	if err := db.Delete(&user).Error; err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}
