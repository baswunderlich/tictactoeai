package ai

import (
	"math/rand"
	"tictactoeai/tictactoe"
)

type RandomAI struct{}

func (r *RandomAI) GetMove(game tictactoe.Game) Coordinate {
	possibleMoves := GetPossibleMoves(game)
	return possibleMoves[rand.Intn(len(possibleMoves))]
}
