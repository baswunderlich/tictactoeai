package ai_test

import (
	"testing"
	"tictactoeai/ai"
	"tictactoeai/tictactoe"
)

func TestSingleCoordinate(t *testing.T) {
	field := [][]tictactoe.FieldValue{
		{2, 1, 2},
		{1, 1, 2},
		{2, 2, 0},
	}
	game := tictactoe.Game{
		Field:         field,
		CurrentPlayer: tictactoe.Cross,
	}
	ai := ai.RandomAI{}
	move := ai.GetMove(game)
	if move.X != 0 || move.Y != 0 {
		t.Errorf("Somethings wrong with the free coordinates")
	}
}
