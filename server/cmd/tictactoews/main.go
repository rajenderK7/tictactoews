package main

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	gamemanager "github.com/rajenderK7/tictactoews/internal/game_manager"
)

var (
	upgrader = websocket.Upgrader{
		// Allow all origins for simplicity.
		CheckOrigin:     func(r *http.Request) bool { return true },
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

var gm = gamemanager.New()

func main() {
	e := echo.New()

	e.GET("/new-game", func(c echo.Context) error {
		// Upgrade HTTP to Websocket.
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		gm.NewGame(ws)
		return c.JSON(http.StatusSwitchingProtocols, "new game initiated")
	})

	e.Start("localhost:4000")
}
