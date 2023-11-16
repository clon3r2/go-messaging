package main

import (
	"main/db"
)

func initialize() {
	db.InitializeDatabase()
	db.MigrateModels()

}

func main() {
	// TODO: check fiber socket
	// TODO: check echo socket
}
