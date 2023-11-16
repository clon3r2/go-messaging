package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"main/db"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	WSConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("WSConn =", WSConn)

}

func initialize() {
	db.InitializeDatabase()
	db.MigrateModels()
}

func main() {

}
