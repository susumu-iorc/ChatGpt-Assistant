package database

import (
	"fmt"
	"log"

	"github.com/susumu-iorc/ChatGPT-Personal-Assistant/app/models"
)

func Seed() {
	db := DbInit()

	users := []models.User{
		{Name: "susumu", Type: "master"},
		{Name: "hal", Type: "normal"},
	}
	result1 := db.Create(&users)

	var user_susumu models.User
	var user_hal models.User
	db.Where("name = ?", "susumu").First(&user_susumu)
	db.Where("name = ?", "hal").First(&user_hal)

	talk_thread := []models.TalkThread{
		{Token: "abc123", UserID: user_susumu.ID},
		{Token: "def456", UserID: user_hal.ID},
	}

	result2 := db.Create(&talk_thread)

	var talk_thread_susumu models.TalkThread
	var talk_thread_hal models.TalkThread

	db.Where("token = ?", "abc123").First(&talk_thread_susumu)
	db.Where("token = ?", "def456").First(&talk_thread_hal)

	talks := []models.Talk{
		{TalkThreadID: talk_thread_susumu.ID, Talker: "user", Content: "こんにちは"},
		{TalkThreadID: talk_thread_susumu.ID, Talker: "bot", Content: "こんにちは"},
		{TalkThreadID: talk_thread_hal.ID, Talker: "user", Content: "おはようございます"},
		{TalkThreadID: talk_thread_susumu.ID, Talker: "user", Content: "今日は暑いですね"},
		{TalkThreadID: talk_thread_hal.ID, Talker: "bot", Content: "おはようございます"},
		{TalkThreadID: talk_thread_susumu.ID, Talker: "bot", Content: "そうですね"},
		{TalkThreadID: talk_thread_susumu.ID, Talker: "user", Content: "なんのご飯を食べましょうか"},
		{TalkThreadID: talk_thread_susumu.ID, Talker: "bot", Content: "カレーがいいですよ"},
		{TalkThreadID: talk_thread_hal.ID, Talker: "user", Content: "今日は何をしましょうか"},
		{TalkThreadID: talk_thread_hal.ID, Talker: "bot", Content: "テニスがいいですよ"},
	}

	result3 := db.Create(&talks)

	if result1.Error != nil || result2.Error != nil || result3.Error != nil {
		log.Fatal(result1.Error)
		log.Fatal(result2.Error)
		log.Fatal(result3.Error)
	}
	fmt.Println("count:", result1.RowsAffected)
	fmt.Println("count:", result2.RowsAffected)
	fmt.Println("count:", result3.RowsAffected)
}
