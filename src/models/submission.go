package models

import (
	"github.com/jinzhu/gorm"
)

type Submission struct {
    gorm.Model
    IsApproved bool `json:"is_approved"`
    Submitters string `json:"submitters"`
    ImgurUrl string `json:"imgur_url"`
}
