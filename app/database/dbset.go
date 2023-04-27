package database

import (
	_ "github.com/go-sql-driver/mysql"
	Const "github.com/susumu-iorc/ChatGPT-Personal-Assistant/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbInit() *gorm.DB {
	dsn := Const.DatabaseUser + ":" + Const.DatabasePass + "@tcp(" + Const.DatabaseHost + ":" + Const.DatabasePort + ")/" + Const.DatabaseName

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
