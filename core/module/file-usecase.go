package module

import (
	"errors"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/fajartd02/mygallery/core/entity"
	"github.com/fajartd02/mygallery/core/repository"
)

var ErrBucketNameNotFound = errors.New("bucket name is not defined")

type FileUsecase interface {
	GetFile(c *gin.Context, filename string) (entity.File, error)
	GetFiles(c *gin.Context) ([]entity.File, error)
	UploadFileToStorage(c *gin.Context, ef entity.File, buffer multipart.File) (entity.File, error)
	DeleteFileFromStorage(c *gin.Context, filename string) error
}

type fileUsecase struct {
	fileRepo   repository.FileRepository
	reportRepo repository.ReportRepository
}

func NewFileUseCase(filerepo repository.FileRepository, reportrepo repository.ReportRepository) FileUsecase {
	return &fileUsecase{
		filerepo,
		reportrepo,
	}
}

func (fuc *fileUsecase) GetFile(c *gin.Context, filename string) (entity.File, error) {
	data, err := fuc.fileRepo.FindSingle(c, filename)
	if err != nil {
		if errors.Is(err, repository.ErrRecordFileNotFound) {
			return entity.File{}, err
		}
		return entity.File{}, err
	}
	return data, nil
}

func (fuc *fileUsecase) GetFiles(c *gin.Context) ([]entity.File, error) {
	data, err := fuc.fileRepo.FindAll(c)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (fuc *fileUsecase) UploadFileToStorage(c *gin.Context, ef entity.File, buffer multipart.File) (entity.File, error) {

	hostName, ok := c.MustGet("host").(string)
	if !ok {
		return entity.File{}, errors.New("failed to parse hostname to string")
	}

	newID, err := uuid.NewUUID()
	if err != nil {
		return entity.File{}, err
	}

	dt := time.Now()
	currYear := strconv.Itoa(dt.Year())
	currMonth := strconv.Itoa(int(dt.Month()))
	currDay := strconv.Itoa(dt.Day())
	date := currYear + "/" + currMonth + "/" + currDay

	fileDTO := &entity.File{
		ID:           newID,
		Filename:     newID.String() + ef.Extension,
		OriginalName: ef.OriginalName,
		Mimetype:     ef.Mimetype,
		Url:          hostName + "/file/static/" + date + "/" + newID.String() + ef.Extension,
		Size:         ef.Size,
		Extension:    ef.Extension,
	}

	dir := "images" + "/" + date
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	dst, err := os.Create(dir + "/" + fileDTO.Filename)
	if err != nil {
		log.Println("error create file", err)
		return entity.File{}, err
	}

	if _, err = io.Copy(dst, buffer); err != nil {
		log.Println("error copy", err)
		return entity.File{}, err
	}

	if ef, err = fuc.fileRepo.Create(c, fileDTO); err != nil {
		log.Println("error create file", err)
		return entity.File{}, err
	}

	return ef, nil
}

func (fuc *fileUsecase) DeleteFileFromStorage(c *gin.Context, filename string) error {
	err := fuc.fileRepo.Delete(c, filename)
	if err != nil {
		return err
	}

	return nil
}
