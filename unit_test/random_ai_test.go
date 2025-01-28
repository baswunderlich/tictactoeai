package unit_test

import (
	"testing"
	"tictactoeai/ai"
	"tictactoeai/tictactoe"
)

func TestSingleCoordinate(t *testing.T) {
	field := [][]tictactoe.FieldValue{
		{0, 1, 2},
		{1, 1, 2},
		{2, 2, 1},
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

func TestNumberOfOneAvailableCoordinate(t *testing.T) {
	field := [][]tictactoe.FieldValue{
		{2, 1, 2},
		{1, 1, 2},
		{2, 2, 0},
	}
	game := tictactoe.Game{
		Field:         field,
		CurrentPlayer: tictactoe.Cross,
	}
	moves := ai.GetPossibleMoves(game)
	if len(moves) != 1 {
		t.Errorf("There were %d moves seen as possible although there is just one", len(moves))
	}
}

func TestNumberOfMultipleAvailableCoordinates(t *testing.T) {
	field := [][]tictactoe.FieldValue{
		{2, 1, 2},
		{1, 0, 2},
		{0, 2, 0},
	}
	game := tictactoe.Game{
		Field:         field,
		CurrentPlayer: tictactoe.Cross,
	}
	moves := ai.GetPossibleMoves(game)
	if len(moves) != 3 {
		t.Errorf("There were %d moves seen as possible although there are three", len(moves))
	}

}
