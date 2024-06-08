package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{
		// Allow all origins for simplicity.
		CheckOrigin:     func(r *http.Request) bool { return true },
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func main() {
	e := echo.New()

	e.GET("/ws", func(c echo.Context) error {
		// Upgrade HTTP to Websocket.
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		fmt.Println("New ws connection")
		// Echo back the messages to the client.
		for {
			msgType, msg, err := ws.ReadMessage()
			if err != nil {
				// Connection might have been closed.
				break
			}

			err = ws.WriteMessage(msgType, msg)
			if err != nil {
				break
			}
		}
		return nil
	})

	e.Start(":4000")
}
