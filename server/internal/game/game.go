package game

import (
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/rajenderK7/tictacgo"
)

type move struct {
	player        *websocket.Conn
	markAndCoords []byte
}

type Game struct {
	game    *tictacgo.Game
	Player1 *websocket.Conn
	Player2 *websocket.Conn
	MovesCh chan move
}

func NewGame(player1, player2 *websocket.Conn) *Game {
	return &Game{
		game:    tictacgo.New(3),
		Player1: player1,
		Player2: player2,
		MovesCh: make(chan move, 10),
	}
}

// Start will start listening to the websocket connections.
func (g *Game) Start() error {
	go readMessages(g.Player1, g.MovesCh)
	go readMessages(g.Player2, g.MovesCh)

	for move := range g.MovesCh {
		markAndCoords := strings.Split(string(move.markAndCoords), "")
		mark := markAndCoords[0]
		row, err := strconv.ParseInt(markAndCoords[0], 10, 32)
		if err != nil {
			continue
		}
		col, err := strconv.ParseInt(markAndCoords[1], 10, 32)
		if err != nil {
			continue
		}
		res, err := g.game.Play(mark, int(row), int(col))
		if err != nil {
			continue
		}
		if len(res.Winner) > 0 {
			winner := WINNER + " " + res.Winner
			g.Player1.WriteMessage(websocket.TextMessage, []byte(winner))
			g.Player2.WriteMessage(websocket.TextMessage, []byte(winner))
			break
		} else if res.IsDraw {
			g.Player1.WriteMessage(websocket.TextMessage, []byte(DRAW))
			g.Player2.WriteMessage(websocket.TextMessage, []byte(DRAW))
			break
		} else if move.player == g.Player1 {
			g.Player2.WriteMessage(websocket.TextMessage, move.markAndCoords)
		} else {
			g.Player1.WriteMessage(websocket.TextMessage, move.markAndCoords)
		}
	}
	return nil
}

func readMessages(ws *websocket.Conn, movesCh chan move) {
	defer ws.Close()
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			return
		}

		movesCh <- move{
			player:        ws,
			markAndCoords: msg,
		}
	}
}
