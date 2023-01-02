package entity

import (
	"gorm.io/gorm"
)

type UserVote struct {
	gorm.Model
	UserID   uint
	ReportID uint
	IsUpVote *bool `json:"isUpVote,omitempty"`
}

type UserVoteInput struct {
	IsUpVote *bool `json:"isUpVote,omitempty"`
}

type Votes struct {
	UpVotes   int64 `json:"upVotes"`
	DownVotes int64 `json:"downVotes"`
}
