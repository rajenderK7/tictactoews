package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	gamemanager "github.com/rajenderK7/tictactoews/internal/game_manager"
)

const DEFAULT_PORT = ":4000"

var (
	upgrader = websocket.Upgrader{
		// Allow all origins for simplicity.
		CheckOrigin:     func(r *http.Request) bool { return true },
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

var gm = gamemanager.New()

func getPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}
	return DEFAULT_PORT
}

func main() {
	e := echo.New()

	e.GET("/new-game", func(c echo.Context) error {
		// Upgrade HTTP to Websocket.
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		gm.NewGame(ws)
		return nil
	})

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Start server
	go func() {
		if err := e.Start(getPort()); err != nil && err == http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
