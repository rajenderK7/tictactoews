package game

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type Player struct {
	ws   *websocket.Conn
	Mark string
}

func NewPlayer(ws *websocket.Conn, mark string) *Player {
	return &Player{
		ws:   ws,
		Mark: mark,
	}
}

func (p *Player) SendMessage(msg *Message) {
	resp, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("marshal err: ", err.Error())
		return
	}
	p.ws.WriteMessage(websocket.TextMessage, resp)
}

func (p *Player) CloseWS() {
	p.ws.Close()
}
