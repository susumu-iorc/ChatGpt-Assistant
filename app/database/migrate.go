package database

import "github.com/susumu-iorc/ChatGPT-Personal-Assistant/app/models"

func Migrate() {
	db := DbInit()
	err3 := db.AutoMigrate(&models.TalkThread{})
	err4 := db.AutoMigrate(&models.Client{})
	err := db.AutoMigrate(&models.User{})
	err2 := db.AutoMigrate(&models.Talk{})
	if err != nil || err2 != nil || err3 != nil || err4 != nil {
		panic("failed to migrate database")
	}
}
