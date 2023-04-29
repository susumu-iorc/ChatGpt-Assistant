package models

import "gorm.io/gorm"

type Talk struct {
	gorm.Model

	TalkThreadID uint
	Talker       string
	Content      string
}
