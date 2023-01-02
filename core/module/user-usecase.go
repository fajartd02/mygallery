package module

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/fajartd02/mygallery/core/entity"
	"github.com/fajartd02/mygallery/core/repository"
)

type UserUsecase interface {
	GetUsers(c *gin.Context) ([]entity.User, error)
	GetUser(c *gin.Context) (entity.User, error)
	Login(c *gin.Context, loginInput entity.LoginInput) (entity.Token, error)
	CreateUser(c *gin.Context, user entity.User) error
	UpdateUser(c *gin.Context, user entity.User) error
	DeleteUser(c *gin.Context) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

// NewUserUsecase use for initiate new user usecase
func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo}
}

func (em *userUsecase) GetUsers(c *gin.Context) ([]entity.User, error) {
	data, err := em.userRepo.FindAll(c)
	if err != nil {
		if errors.Is(err, repository.ErrRecordUserNotFound) {
			return nil, fmt.Errorf("%w.", ErrUserNotFound)
		}
		return nil, fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}
	return data, nil
}

func (em *userUsecase) GetUser(c *gin.Context) (entity.User, error) {
	data, err := em.userRepo.FindSingle(c, 1)
	if err != nil {
		if errors.Is(err, repository.ErrRecordUserNotFound) {
			return entity.User{}, fmt.Errorf("%w.", ErrUserNotFound)
		}
		return entity.User{}, fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}
	return data, nil
}

func (em *userUsecase) Login(c *gin.Context, loginInput entity.LoginInput) (entity.Token, error) {

	authUser, err := em.userRepo.FindSingleByEmail(c, loginInput.Email)
	if err != nil {
		return entity.Token{}, fmt.Errorf("email not exists")
	}

	check := CheckPasswordHash(loginInput.Password, authUser.Password)
	if !check {
		return entity.Token{}, fmt.Errorf("username or password is incorrect")
	}

	validToken, err := GenerateJWT(loginInput.Email, authUser.ID)
	if err != nil {
		return entity.Token{}, fmt.Errorf("invalid token")
	}

	var token entity.Token
	token.UserID = authUser.ID
	token.Email = authUser.Email
	token.TokenString = validToken
	return token, nil
}

func (em *userUsecase) CreateUser(c *gin.Context, user entity.User) error {
	_, err := em.userRepo.FindSingleByEmail(c, user.Email)
	if err != repository.ErrRecordUserNotFound {
		return fmt.Errorf("email is already exists")
	}

	hashedPass, err := GeneratehashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("error in password hash")
	}
	user.Password = hashedPass

	err = em.userRepo.Create(c, user)
	if err != nil {
		if errors.Is(err, repository.ErrRecordUserNotFound) {
			return fmt.Errorf("%w.", ErrUserNotFound)
		}
		return fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}
	return nil
}

func (em *userUsecase) UpdateUser(c *gin.Context, user entity.User) error {
	userByEmail, err := em.userRepo.FindSingleByEmail(c, user.Email)
	if err != nil {
		return fmt.Errorf("email not exists")
	}

	check := CheckPasswordHash(userByEmail.Password, user.Password)
	if !check {
		hashedPass, err := GeneratehashPassword(user.Password)
		if err != nil {
			return fmt.Errorf("error in password hash")
		}
		user.Password = hashedPass
	}

	err = em.userRepo.Update(c, user)
	if err != nil {
		if errors.Is(err, repository.ErrRecordUserNotFound) {
			return fmt.Errorf("%w.", ErrUserNotFound)
		}
		return fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}
	return nil
}

func (em *userUsecase) DeleteUser(c *gin.Context) error {
	err := em.userRepo.Delete(c)
	if err != nil {
		if errors.Is(err, repository.ErrRecordUserNotFound) {
			return fmt.Errorf("%w.", ErrUserNotFound)
		}
		return fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}
	return nil
}
