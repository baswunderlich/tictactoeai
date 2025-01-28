package ai

import (
	"tictactoeai/tictactoe"
)

type Brain interface {
	GetMove(tictactoe.Game) Coordinate
}
