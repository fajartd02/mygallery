package module

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/fajartd02/mygallery/backend/core/entity"
	"github.com/fajartd02/mygallery/backend/core/repository"
)

var ErrMemoryNotFound = errors.New("memory error: ")

type MemoryUseCase interface {
	CreateMemory(c *gin.Context, memory entity.Memory) (entity.Memory, error)
	DeleteMemory(c *gin.Context) error
	FindAll(c *gin.Context) ([]entity.Memory, error)
	FindByID(c *gin.Context) (entity.Memory, error)
	UpdateMemory(c *gin.Context, memory entity.Memory) error
	List(c *gin.Context, memoryList entity.MemoryListRequest) ([]entity.Memory, error)
}

type memoryUseCase struct {
	memoRepo repository.MemoryRepository
}

func NewMemoryUseCase(repo repository.MemoryRepository) MemoryUseCase {
	return &memoryUseCase{repo}
}

func (em memoryUseCase) CreateMemory(c *gin.Context, memory entity.Memory) (entity.Memory, error) {
	memory, err := em.memoRepo.Create(c, memory)
	if err != nil {
		if errors.Is(err, repository.ErrRecordMemoryNotFound) {
			return entity.Memory{}, fmt.Errorf("%w.", ErrMemoryNotFound)
		}
		return entity.Memory{}, fmt.Errorf("%w: %v", ErrMemoryNotFound, err)
	}
	return memory, err
}

func (em memoryUseCase) DeleteMemory(c *gin.Context) error {
	err := em.memoRepo.Delete(c)
	if err != nil {
		if errors.Is(err, repository.ErrRecordMemoryNotFound) {
			return fmt.Errorf("%w.", ErrMemoryNotFound)
		}
		return fmt.Errorf("%w: %v", ErrMemoryNotFound, err)
	}
	return nil
}

func (em memoryUseCase) FindAll(c *gin.Context) ([]entity.Memory, error) {
	memories, err := em.memoRepo.FindAll(c)
	if err != nil {
		if errors.Is(err, repository.ErrRecordMemoryNotFound) {
			return nil, fmt.Errorf("%w.", ErrMemoryNotFound)
		}
		return nil, fmt.Errorf("%w: %v", ErrMemoryNotFound, err)
	}
	return memories, nil
}

func (em memoryUseCase) FindByID(c *gin.Context) (entity.Memory, error) {
	memory, err := em.memoRepo.FindByID(c)
	if err != nil {
		if errors.Is(err, repository.ErrRecordMemoryNotFound) {
			return entity.Memory{}, fmt.Errorf("%w.", ErrMemoryNotFound)
		}
		return entity.Memory{}, fmt.Errorf("%w: %v", ErrMemoryNotFound, err)
	}
	return memory, nil
}

func (em memoryUseCase) List(c *gin.Context, memoryList entity.MemoryListRequest) ([]entity.Memory, error) {
	if memoryList.Filter.By == "" {
		memoryList.Filter.By = "tag"
	}

	if memoryList.Sort.By == "" {
		memoryList.Sort.By = "tag"
	}

	memories, err := em.memoRepo.List(c, memoryList)
	if err != nil {
		if errors.Is(err, repository.ErrRecordMemoryNotFound) {
			return nil, fmt.Errorf("%w.", ErrMemoryNotFound)
		}
		return nil, fmt.Errorf("%w: %v", ErrMemoryNotFound, err)
	}
	return memories, nil
}

func (em memoryUseCase) UpdateMemory(c *gin.Context, memory entity.Memory) error {
	err := em.memoRepo.Update(c, memory)
	if err != nil {
		if errors.Is(err, repository.ErrRecordMemoryNotFound) {
			return fmt.Errorf("%w.", ErrMemoryNotFound)
		}
		return fmt.Errorf("%w: %v", ErrMemoryNotFound, err)
	}
	return nil
}
