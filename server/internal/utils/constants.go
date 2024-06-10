package utils

const (
	// Game states
	OK                        = "OK"
	START                     = "START"
	DRAW                      = "DRAW"
	GAME_OVER                 = "GAME OVER"
	WINNER_TITLE              = "WINNER"
	OPPONENT_DISCONNECTED_WIN = "Opponent disconnected. You win!"

	// Player message types
	INIT_GAME    = "init_game"
	MOVE         = "move"
	DISCONNECTED = "disconnected"
	ERROR        = "error"
	WAITING      = "WAITING"

	// Player statuses
	WAITING_FOR_OPPONENT = "WAITING FOR OPPONENT"
)
