package main

import (
	"github.com/susumu-iorc/ChatGPT-Personal-Assistant/app/database"
	"github.com/susumu-iorc/ChatGPT-Personal-Assistant/app/server"
)

func main() {
	// config.Const()
	// fmt.Println(Const.DatabaseUser)
	// fmt.Println(Const.DatabasePass)
	// fmt.Println(Const.Databasepass)
	db := database.Migrate
	// database.Seed()
	db()

	server.Router()
}
