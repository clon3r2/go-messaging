package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main/db"
	"main/server"
	"net/http"
	"time"
)

func init() {
	db.InitializeDatabase()
	db.MigrateModels()
	server.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
}

func main() {
	// TODO: check fiber socket
	// TODO: check echo socket
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/socket/", SocketHandler)
	e.GET("/socket/ping/", Pinger)
	e.Logger.Fatal(e.Start(":8000"))
}

func Pinger(c echo.Context) error {
	WSConn, err := server.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Error(err)
	}
	defer func(WSConn *websocket.Conn) {
		err := WSConn.Close()
		if err != nil {

		}
	}(WSConn)

	err = WSConn.WriteControl(websocket.PongMessage, []byte("pong"), time.Now().Add(time.Hour))
	if err != nil {
		c.Logger().Error(err)
	}
	return nil
}

func SocketHandler(c echo.Context) error {
	WSConn, err := server.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Error(err)
	}
	defer func(WSConn *websocket.Conn) {
		err := WSConn.Close()
		if err != nil {

		}
	}(WSConn)
	i := 0
	for {
		err = WSConn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("-`----1---->%v<--------", i)))
		if err != nil {
			c.Logger().Error(err)
		}
		i++
		time.Sleep(time.Second)

	}
}
