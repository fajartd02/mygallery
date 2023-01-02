package entity

import (
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	UserVotes   []UserVote `json:"userVotes"`
	Comments    []Comment  `json:"comments"`
	UserID      uint
	EventTitle  string  `json:"eventTitle"`
	Category    uint    `json:"category"`
	ImageUrl    string  `json:"imageUrl"`
	Description string  `json:"description"`
	Longitude   string  `json:"longitude"`
	Latitude    string  `json:"latitude"`
	Radius      float64 `json:"radius"`
}

type ReportResp struct {
	gorm.Model
	UserVotes   []UserVote `json:"userVotes"`
	Comments    []Comment  `json:"comments"`
	UserID      uint
	EventTitle  string  `json:"eventTitle"`
	Category    uint    `json:"category"`
	ImageUrl    string  `json:"imageUrl"`
	Description string  `json:"description"`
	Longitude   string  `json:"longitude"`
	Latitude    string  `json:"latitude"`
	Radius      float64 `json:"radius"`
	UpVotes     int64   `json:"upVotes"`
	DownVotes   int64   `json:"downVotes"`
}

type ReportInput struct {
	UserVotes   []UserVoteInput `json:"userVotes"`
	Comments    []CommentInput  `json:"comments"`
	EventTitle  string          `json:"eventTitle"`
	Category    uint            `json:"category"`
	ImageUrl    string          `json:"imageUrl"`
	Description string          `json:"description"`
	Longitude   string          `json:"longitude"`
	Latitude    string          `json:"latitude"`
	Radius      float64         `json:"radius"`
}
