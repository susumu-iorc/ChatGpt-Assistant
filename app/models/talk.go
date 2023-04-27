package models

import "gorm.io/gorm"

type Talk struct {
	gorm.Model
	TalkThread    string
	Content       string
	PrevContentId int
}
