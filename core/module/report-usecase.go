package module

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/fajartd02/mygallery/core/entity"
	"github.com/fajartd02/mygallery/core/repository"
)

type ReportUsecase interface {
	GetReports(c *gin.Context) ([]entity.Report, error)
	GetReport(c *gin.Context, reportId uint) (entity.ReportResp, error)
	CreateReport(c *gin.Context, user entity.Report) error
	UpdateReport(c *gin.Context, user entity.Report) error
	DeleteReport(c *gin.Context) error
}

type reportUsecase struct {
	reportRepo repository.ReportRepository
	userRepo   repository.UserRepository
}

func NewReportUseCase(reportrepo repository.ReportRepository, userrepo repository.UserRepository) ReportUsecase {
	return &reportUsecase{
		reportrepo,
		userrepo,
	}
}

func (r reportUsecase) GetReports(c *gin.Context) ([]entity.Report, error) {
	data, err := r.reportRepo.FindAll(c)
	if err != nil {
		if errors.Is(err, repository.ErrRecordUserNotFound) {
			return []entity.Report{}, fmt.Errorf("%w.", ErrUserNotFound)
		}
		return []entity.Report{}, fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}
	return data, nil
}

func (r reportUsecase) GetReport(c *gin.Context, reportId uint) (entity.ReportResp, error) {
	data, err := r.reportRepo.FindSingle(c, reportId)
	if err != nil {
		if errors.Is(err, repository.ErrRecordUserNotFound) {
			return entity.ReportResp{}, fmt.Errorf("%w.", ErrRecordReportNotFound)
		}
		return entity.ReportResp{}, fmt.Errorf("%w: %v", ErrRecordReportNotFound, err)
	}

	return data, nil
}

func (r reportUsecase) CreateReport(c *gin.Context, report entity.Report) error {
	report, err := r.reportRepo.Create(c, report)
	if err != nil {
		if errors.Is(err, repository.ErrRecordReportNotFound) {
			return fmt.Errorf("%w.", ErrRecordReportNotFound)
		}
		return fmt.Errorf("%w: %v", ErrRecordReportNotFound, err)
	}
	return nil
}

func (r reportUsecase) UpdateReport(c *gin.Context, user entity.Report) error {
	//TODO implement me
	panic("implement me")
}

func (r reportUsecase) DeleteReport(c *gin.Context) error {
	//TODO implement me
	panic("implement me")
}
