package models

import (
	"github.com/jinzhu/gorm"
)

type Activity struct {
    gorm.Model
    CategoryId int `json:"category_id"`
    Name string `json:"name"`
}
