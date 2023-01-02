package reportrepository

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fajartd02/mygallery/core/entity"
	repository_intf "github.com/fajartd02/mygallery/core/repository"
)

var ErrDBNotFound = errors.New("repo: db not found")

type repository struct {
}

func (r repository) FindAll(c *gin.Context) ([]entity.Report, error) {
	var reports []entity.Report

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return nil, errors.New("failed to parse db to gorm")
	}

	if err := db.Model(&[]entity.Report{}).Preload("Comments").Preload("UserVotes").Find(&reports).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrRecordReportNotFound
		}
		return nil, err
	}

	return reports, nil
}

func (r repository) FindSingle(c *gin.Context, reportID uint) (entity.ReportResp, error) {
	report := entity.Report{}

	var userVotes []entity.UserVote
	var downvotes int64 = 0
	var upvotes int64 = 0

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return entity.ReportResp{}, errors.New("failed to parse db to gorm")
	}

	db.Raw("SELECT * FROM user_votes WHERE report_id = ?", reportID).Scan(&userVotes)

	if userVotes == nil {
		return entity.ReportResp{}, errors.New("user votes not found")
	}

	for _, vote := range userVotes {
		if vote.IsUpVote != nil {
			if *vote.IsUpVote {
				upvotes += 1
			} else {
				downvotes += 1
			}
		}
	}

	var votes = entity.Votes{
		UpVotes:   upvotes,
		DownVotes: downvotes,
	}

	if err := db.Model(&entity.Report{}).Preload("Comments").Preload("UserVotes").Where("id = ?", reportID).First(&report).Error; err != nil {
		return entity.ReportResp{}, repository_intf.ErrRecordUserNotFound
	}

	resp := entity.ReportResp{
		Model:       report.Model,
		UserVotes:   report.UserVotes,
		Comments:    report.Comments,
		UserID:      report.UserID,
		EventTitle:  report.EventTitle,
		Category:    report.Category,
		ImageUrl:    report.ImageUrl,
		Description: report.Description,
		Longitude:   report.Longitude,
		Latitude:    report.Latitude,
		Radius:      report.Radius,
		UpVotes:     votes.UpVotes,
		DownVotes:   votes.DownVotes,
	}

	return resp, nil
}

func (r repository) Create(c *gin.Context, report entity.Report) (entity.Report, error) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return entity.Report{}, errors.New("failed to parse db to gorm")
	}

	if err := db.Create(&report).Error; err != nil {
		return entity.Report{}, errors.New("failed to create report")
	}

	return report, nil
}

func (r repository) Delete(c *gin.Context, filename string) error {
	//TODO implement me
	panic("implement me")
}

func New() repository_intf.ReportRepository {
	return &repository{}
}
