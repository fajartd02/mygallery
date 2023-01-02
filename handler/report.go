package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/fajartd02/mygallery/core/entity"
	"github.com/fajartd02/mygallery/core/module"
)

type ReportHandler struct {
	reportUc module.ReportUsecase
}

func NewReportHandler(reportUc module.ReportUsecase) *ReportHandler {
	return &ReportHandler{
		reportUc,
	}
}

func (hdl *ReportHandler) CreateReport(c *gin.Context) {
	var input entity.ReportInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	userIdHeader := c.Request.Header["User-Id"]
	if len(userIdHeader) == 0 {
		c.AbortWithStatusJSON(500, gin.H{"message": "userID not found in the header!"})
		return
	}

	userID, err := strconv.ParseUint(userIdHeader[0], 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	// Create Report
	report := entity.Report{
		UserID:      uint(userID),
		EventTitle:  input.EventTitle,
		Category:    input.Category,
		ImageUrl:    input.ImageUrl,
		Description: input.Description,
		Latitude:    input.Latitude,
		Longitude:   input.Longitude,
		Radius:      input.Radius,
	}

	err = hdl.reportUc.CreateReport(c, report)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "report succesfully created"})
}

func (hdl *ReportHandler) GetReports(c *gin.Context) {
	reports, err := hdl.reportUc.GetReports(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reports})
}

func (hdl *ReportHandler) GetReport(c *gin.Context) {
	reportId, err := strconv.ParseUint(c.Param("reportId"), 10, 32)

	report, err := hdl.reportUc.GetReport(c, uint(reportId))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": report})
}
