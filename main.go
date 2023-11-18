package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main/db"
	"main/server"
	"net/http"
)

func MakeNewServer() Server {
	return Server{
		pool: make(map[*websocket.Conn]bool),
	}
}

type Server struct {
	pool map[*websocket.Conn]bool
}

func (s *Server) NewConnectionHandler(c echo.Context) error {
	WSConn, err := server.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Error(err)
	}
	s.pool[WSConn] = true
	fmt.Printf("new connection added to pool. total number of pool connections: %v\n\n\n", len(s.pool))
	go s.ReadLoop(WSConn)
	return nil
}

func (s *Server) ReadLoop(c *websocket.Conn) {
	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("new msg => %s", msg)
		go s.broadCast(msg)
	}
}

func (s *Server) broadCast(msg []byte) {
	for con := range s.pool {
		if s.pool[con] {
			err := con.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				fmt.Printf("error writing msg to conn %v", con)
				continue
			}
		}
	}
}

func init() {
	db.InitializeDatabase()
	db.MigrateModels()
	server.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
}

func main() {
	baseServer := MakeNewServer()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/socket/", baseServer.NewConnectionHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
