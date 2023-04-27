package database

import "github.com/susumu-iorc/ChatGPT-Personal-Assistant/app/models"

func Migrate() {
	db := DbInit()
	err := db.AutoMigrate(&models.User{})
	err2 := db.AutoMigrate(&models.Talk{})
	err3 := db.AutoMigrate(&models.TalkThread{})
	if err != nil || err2 != nil || err3 != nil {
		panic("failed to migrate database")
	}
}
