package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Conn *gorm.DB
var er error

func InitializeDatabase() {
	Conn, er = gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if er != nil {
		log.Fatal(er)
	}
}

func MigrateModels() {
	er = Conn.AutoMigrate(&User{})
	if er != nil {
		log.Fatal(er)
	}
	er = Conn.AutoMigrate(&Chat{})
	if er != nil {
		log.Fatal(er)
	}
	er = Conn.AutoMigrate(&Message{})
	if er != nil {
		log.Fatal(er)
	}
}
