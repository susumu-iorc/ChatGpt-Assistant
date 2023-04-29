package Const

import (
	"os"

	"github.com/joho/godotenv"
)

func GetConst(s string) string {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}
	return os.Getenv(s)
}

var DatabaseUser = GetConst("DATABASE_USER")
var DatabasePass = GetConst("DATABASE_PASS")
var DatabaseHost = GetConst("DATABASE_HOST")
var DatabasePort = GetConst("DATABASE_PORT")
var DatabaseName = GetConst("DATABASE_NAME")

var OpenAIApiKey = GetConst("OPENAI_API_KEY")
