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
	//TODO: add authentication check here before making socket connection
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
		fmt.Printf("\n new msg => %s\n", msg)
		go s.broadCast(msg)
	}
}

func (s *Server) broadCast(msg []byte) {
	fmt.Println("\n#log all conns: ")
	for con := range s.pool {
		fmt.Printf("addr => %s    status => %v\n", con.RemoteAddr(), s.pool[con])
	}
	fmt.Println("# end log\n")
	for con := range s.pool {
		if s.pool[con] {
			err := con.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				fmt.Printf("error writing msg to conn %v, connection closed.\n", con.RemoteAddr())
				s.pool[con] = false
				continue
			}
			fmt.Printf("successfully sent '%s' to %v\n", msg, con.RemoteAddr())
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/socket/", baseServer.NewConnectionHandler)
	e.POST("/signup/", SignUp)
	e.POST("/login/", Login)
	e.Logger.Fatal(e.Start(":8080"))
}

func Login(c echo.Context) error {
	// TODO: login does not work properly
	type loginData struct {
		Username string
		Password string
	}
	var payload loginData
	err := c.Bind(&payload)
	if err != nil {
		fmt.Println("api err => ", err)
		return c.JSON(400, err.Error())
	}
	fmt.Printf("\n\ngot payload => %+v\n\n", payload)
	user, err := server.UserLogin(payload.Username, payload.Password)
	if err != nil {
		fmt.Println("service err ==> ", err)
		return c.JSON(400, err.Error())
	}

	return c.JSON(200, user)
}

func SignUp(c echo.Context) error {
	// tOdo: add data validation
	var user db.User
	err := c.Bind(&user)
	if err != nil {
		fmt.Println("user payload====>")
		return echo.NewHTTPError(400, err)
	}
	result := server.UserSignUp(&user)
	if result.Error != nil {
		fmt.Printf("res ==> %+v", result)
		fmt.Println("db err => ", result.Error)
		return echo.NewHTTPError(400, result.Error)
	}
	fmt.Printf("\n\n\nresult == > %+v", result)
	return c.JSON(200, user)
}
