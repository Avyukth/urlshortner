package model

import "github.com/jinzhu/gorm"

type (
	UrlModel struct {
		gorm.Model
		UrlHashId string `gorm:"primary_key"`
		Url       string
		Shorten   string
		Hits      int
	}
)
