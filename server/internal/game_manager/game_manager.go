package gamemanager

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/rajenderK7/tictactoews/internal/game"
	"github.com/rajenderK7/tictactoews/internal/utils"
)

const (
	MarkX = "X"
	MarkO = "O"
)

type GameManager struct {
	games         []*game.Game
	mu            *sync.RWMutex
	waitingPlayer *game.Player
}

func New() *GameManager {
	return &GameManager{
		games: make([]*game.Game, 0),
		mu:    &sync.RWMutex{},
	}
}

func (gm *GameManager) NewGame(ws *websocket.Conn) {
	gm.mu.Lock()
	defer gm.mu.Unlock()

	if gm.waitingPlayer == nil {
		gm.waitingPlayer = game.NewPlayer(ws, MarkX)
		gm.waitingPlayer.SendMessage(game.NewMessage(nil, utils.WAITING, utils.WAITING_FOR_OPPONENT))
	} else {
		newPlayer := game.NewPlayer(ws, MarkO)
		newGame := game.NewGame(gm.waitingPlayer, newPlayer)
		gm.waitingPlayer = nil
		// gm.addGame(newGame)
		newGame.Player1.SendMessage(game.NewMessage(nil, utils.START, newGame.Player1.Mark))
		newGame.Player2.SendMessage(game.NewMessage(nil, utils.START, newGame.Player2.Mark))
		go newGame.Start()
	}
}

// func (gm *GameManager) addGame(game *game.Game) {
// 	gm.games = append(gm.games, game)
// }
