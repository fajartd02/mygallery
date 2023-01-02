package repository

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/fajartd02/mygallery/core/entity"
)

var ErrRecordReportNotFound = errors.New("record not found")

type ReportRepository interface {
	FindAll(c *gin.Context) ([]entity.Report, error)
	FindSingle(c *gin.Context, reportID uint) (entity.ReportResp, error)
	Create(c *gin.Context, report entity.Report) (entity.Report, error)
	Delete(c *gin.Context, filename string) error
}
