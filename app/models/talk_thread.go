package models

import "gorm.io/gorm"

type TalkThread struct {
	gorm.Model

	Token  string
	UserID uint
}
