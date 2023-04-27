package models

import "gorm.io/gorm"

type TalkThread struct {
	gorm.Model
	TalkToken    string
	LatestTalkId int
}
