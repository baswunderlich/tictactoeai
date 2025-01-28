package unit_test

import (
	"testing"
	"tictactoeai/tictactoe"
)

func TestStatusOngoing(t *testing.T) {
	got := [][]tictactoe.FieldValue{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	game := tictactoe.Game{Field: got}
	status := game.GetStatus()
	if status != tictactoe.GameStatus(tictactoe.OnGoing) {
		t.Errorf("Game was not recognized as OnGoing %d", status)
	}
}

func TestStatusTie(t *testing.T) {
	got := [][]tictactoe.FieldValue{
		{2, 1, 2},
		{1, 1, 2},
		{2, 2, 1},
	}
	game := tictactoe.Game{Field: got}
	status := game.GetStatus()
	if status != tictactoe.GameStatus(tictactoe.Tie) {
		t.Errorf("Game was not recognized as Tie: %d", status)
	}
}

func TestStatusVictoryHorizontal(t *testing.T) {
	got := [][]tictactoe.FieldValue{
		{2, 1, 2},
		{1, 1, 1},
		{2, 2, 1},
	}
	game := tictactoe.Game{Field: got}
	status := game.GetStatus()
	if status != tictactoe.GameStatus(tictactoe.CrossVictory) {
		t.Errorf("Game was not recognized as CrossVictory %d", status)
	}
}

func TestStatusVictoryVertical(t *testing.T) {
	got := [][]tictactoe.FieldValue{
		{1, 1, 2},
		{1, 2, 1},
		{1, 2, 1},
	}
	game := tictactoe.Game{Field: got}
	if game.GetStatus() != tictactoe.GameStatus(tictactoe.CrossVictory) {
		t.Errorf("Game was not recognized as CrossVictory")
	}
}

func TestStatusVictoryDiagonal(t *testing.T) {
	got := [][]tictactoe.FieldValue{
		{1, 1, 2},
		{1, 1, 2},
		{2, 2, 1},
	}
	game := tictactoe.Game{Field: got}
	if game.GetStatus() != tictactoe.GameStatus(tictactoe.CrossVictory) {
		t.Errorf("Game was not recognized as CrossVictory")
	}
}
