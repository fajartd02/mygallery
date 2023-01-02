package uservoterepository

import (
	"errors"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fajartd02/mygallery/core/entity"
	repository_intf "github.com/fajartd02/mygallery/core/repository"
)

type repositoryUserVote struct {
}

func New() repository_intf.UserVoteRepository {
	return &repositoryUserVote{}
}

func (r *repositoryUserVote) GetVotesInReport(c *gin.Context) (entity.Votes, error) {
	var userVotes []entity.UserVote
	var downvotes int64 = 0
	var upvotes int64 = 0
	paramsId, err := strconv.ParseInt(c.Params.ByName("reportId"), 32, 32)
	if err != nil {
		return entity.Votes{}, errors.New("failed to convert params")
	}
	reportId := uint(paramsId)

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return entity.Votes{}, errors.New("failed to parse db to gorm")
	}

	db.Raw("SELECT * FROM user_votes WHERE report_id = ?", reportId).Scan(&userVotes)

	if userVotes == nil {
		return entity.Votes{}, errors.New("user votes not found")
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

	return votes, nil
}

func (r *repositoryUserVote) ChooseVotes(c *gin.Context) error {
	userIdHeader := c.Request.Header["User-Id"]
	if len(userIdHeader) == 0 {
		return errors.New("failed to parse userID")
	}

	userId, err := strconv.ParseUint(userIdHeader[0], 10, 32)
	if err != nil {
		return errors.New("failed to parse userID")
	}

	query := c.Query("isVoting")
	isVoting, err := strconv.ParseBool(query)

	if err != nil {
		log.Fatal(err)
	}

	paramsId, err := strconv.ParseInt(c.Params.ByName("reportId"), 32, 32)
	if err != nil {
		return errors.New("failed to convert params")
	}
	reportId := uint(paramsId)

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return errors.New("failed to parse db to gorm")
	}

	userVote := entity.UserVote{}
	dbRresult := db.Where("user_id = ? AND report_id = ?", userId, reportId).First(&userVote)

	// when user never vote the report that we addresed
	if errors.Is(dbRresult.Error, gorm.ErrRecordNotFound) {
		vote := entity.UserVote{
			UserID:   uint(userId),
			ReportID: reportId,
			IsUpVote: &isVoting,
		}
		if err := db.Create(&vote).Error; err != nil {
			return errors.New("failed to create vote")
		}
	} else {
		// nil to isVoting
		if userVote.IsUpVote == nil {
			if !isVoting {
				err := db.Raw(` 
				UPDATE user_votes
				SET is_up_vote = FALSE
				WHERE user_id = ?;`, userId).Scan(&userVote).Error

				if err != nil {
					return errors.New("failed to convert nil to isVoting")
				}

			} else {
				err := db.Raw(` 
				UPDATE user_votes
				SET is_up_vote = TRUE
				WHERE user_id = ?;`, userId).Scan(&userVote).Error

				if err != nil {
					return errors.New("failed to convert nil to isVoting")
				}
			}
		} else {
			if *userVote.IsUpVote && isVoting {
				// update uservotes ke nil
				err := db.First(&userVote).Error
				if err != nil {
					return errors.New("failed convert from vote true to nil")
				}

				userVote.IsUpVote = nil

				err = db.Save(&userVote).Error
				if err != nil {
					return errors.New("failed save from vote true to nil")
				}

			} else if *userVote.IsUpVote && !isVoting {
				// update uservotes ke false
				err := db.First(&userVote).Error
				if err != nil {
					return errors.New("failed convert from vote true to nil")
				}

				*userVote.IsUpVote = false

				err = db.Save(&userVote).Error
				if err != nil {
					return errors.New("failed save from vote true to nil")
				}

			} else if !*userVote.IsUpVote && !isVoting {
				// update uservotes jadi nill
				err := db.First(&userVote).Error
				if err != nil {
					return errors.New("failed convert from vote false to nil")
				}

				userVote.IsUpVote = nil

				err = db.Save(&userVote).Error
				if err != nil {
					return errors.New("failed save from vote true to nil")
				}

			} else if !*userVote.IsUpVote && isVoting {
				// update uservotes ke true
				err := db.First(&userVote).Error
				if err != nil {
					return errors.New("failed convert from vote false to true")
				}

				*userVote.IsUpVote = true

				db.Save(&userVote)
				if err != nil {
					return errors.New("failed convert from vote true to nil")
				}
			}
		}
	}

	return nil
}
