package game

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/rajenderK7/tictacgo"
	"github.com/rajenderK7/tictactoews/internal/utils"
)

type Message struct {
	Player  *Player `json:"-"`
	Type    string  `json:"type"`
	Payload string  `json:"payload,omitempty"`
}

func NewMessage(player *Player, t string, payload string) *Message {
	return &Message{
		Player:  player,
		Type:    t,
		Payload: payload,
	}
}

type Game struct {
	game      *tictacgo.Game
	Player1   *Player
	Player2   *Player
	MessageCh chan Message
}

func NewGame(player1, player2 *Player) *Game {
	return &Game{
		game:      tictacgo.New(3),
		Player1:   player1,
		Player2:   player2,
		MessageCh: make(chan Message, 5),
	}
}

// Start will listen to the websocket connections.
// The function returns when one of the players disconnect.
func (g *Game) Start() error {
	go readMessages(g.Player1, g.MessageCh)
	go readMessages(g.Player2, g.MessageCh)
	defer g.endGame()

	for msg := range g.MessageCh {
		switch msg.Type {
		case utils.MOVE:
			{
				coords := strings.Split(msg.Payload, " ")

				row, col, err := parseCoords(coords)
				if err != nil {
					msg.Player.SendMessage(NewMessage(nil, utils.ERROR, err.Error()))
					continue
				}

				res, err := g.game.Play(msg.Player.Mark, row, col)
				if err != nil {
					msg.Player.SendMessage(NewMessage(nil, utils.ERROR, err.Error()))
					continue
				}

				move := msg.Player.Mark + " " + msg.Payload
				message := NewMessage(nil, utils.MOVE, move)
				g.messagePlayers(message)

				if len(res.Winner) > 0 {
					message := NewMessage(nil, utils.GAME_OVER, utils.WINNER_TITLE+" "+res.Winner)
					g.messagePlayers(message)
					return nil
				} else if res.IsDraw {
					message := NewMessage(nil, utils.GAME_OVER, utils.DRAW)
					g.messagePlayers(message)
					return nil
				}
			}

		case utils.DISCONNECTED:
			{
				message := NewMessage(nil, utils.GAME_OVER, utils.OPPONENT_DISCONNECTED_WIN)

				if msg.Player == g.Player1 {
					g.Player2.SendMessage(message)
				} else {
					g.Player1.SendMessage(message)
				}

				return nil
			}
		}
	}
	return nil
}

func (g *Game) messagePlayers(msg *Message) {
	g.Player1.SendMessage(msg)
	g.Player2.SendMessage(msg)
}

func (g *Game) endGame() {
	if g.Player1 != nil {
		g.Player1.CloseWS()
	}
	if g.Player2 != nil {
		g.Player2.CloseWS()
	}
	<-g.MessageCh
	<-g.MessageCh
	close(g.MessageCh)
}

func readMessages(player *Player, ch chan Message) {
	defer func() {
		if ch != nil {
			ch <- *NewMessage(player, utils.DISCONNECTED, "")
		}
	}()

	for {
		_, msg, err := player.ws.ReadMessage()
		if err != nil {
			return
		}

		var message Message
		err = json.Unmarshal(msg, &message)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// Assigning the player explicitly is required
		// because we are not passing Player from the client.
		message.Player = player
		ch <- message
	}
}

func parseCoords(coords []string) (int, int, error) {
	row, err := strconv.ParseInt(coords[0], 10, 32)
	if err != nil {
		return -1, -1, err
	}
	col, err := strconv.ParseInt(coords[1], 10, 32)
	if err != nil {
		return -1, -1, err
	}
	return int(row), int(col), err
}
