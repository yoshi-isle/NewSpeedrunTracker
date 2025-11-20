package models

import (
	"github.com/jinzhu/gorm"
)

type Submission struct {
    gorm.Model
    ActivityId int `json:"activity_id"`
    Score int `json:"score"`
    IsActive bool `json:"is_active"`
    IsApproved bool `json:"is_approved"`
    Submitters string `json:"submitters"`
    ImgurUrl string `json:"imgur_url"`
}
