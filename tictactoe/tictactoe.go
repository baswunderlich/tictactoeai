package tictactoe

import (
	"fmt"
)

type FieldValue int

const (
	Cross  FieldValue = 1
	Circle FieldValue = 2
	Empty  FieldValue = 0
)

type GameStatus int

const (
	CircleVictory GameStatus = 1
	CrossVictory  GameStatus = -1
	Tie           GameStatus = 0
	OnGoing       GameStatus = 2
)

type Game struct {
	Field         [][]FieldValue
	CurrentPlayer FieldValue
}

func generate_field() [][]FieldValue {
	field := make([][]FieldValue, 3)
	for i := range field {
		field[i] = make([]FieldValue, 3)
	}
	return field
}

/*
Creates a new Game, circle begins
*/
func CreateGame() Game {
	return Game{
		Field:         generate_field(),
		CurrentPlayer: Circle,
	}
}

func (g *Game) PrintField() {
	for i := range g.Field {
		fmt.Println(g.Field[i])
	}
}

func (g *Game) PlaceCross(x int, y int) error {
	if g.CurrentPlayer == Circle {
		return fmt.Errorf("not that players turn")
	}
	if x < 4 && x > -1 && y < 4 && y > -1 {
		if g.Field[x][y] != Empty {
			return fmt.Errorf("this spot is already in not free")
		}
		g.Field[x][y] = Cross
	}
	g.CurrentPlayer = Circle
	return nil
}

func (g *Game) PlaceCircle(x int, y int) error {
	if g.CurrentPlayer == Cross {
		return fmt.Errorf("not that players turn")
	}
	if x < 4 && x > -1 && y < 4 && y > -1 {
		if g.Field[x][y] != Empty {
			return fmt.Errorf("this spot is already in not free")
		}
		g.Field[x][y] = Circle
	}
	g.CurrentPlayer = Cross
	return nil
}

func (g *Game) GetStatus() GameStatus {

	//Check for circle victory
	if isVictory(g.Field, Circle) {
		return GameStatus(CircleVictory)
	}

	//Check for cross victory
	if isVictory(g.Field, Cross) {
		return GameStatus(CrossVictory)
	}

	//Check for tie
	if isTie(g.Field) {
		return GameStatus(Tie)
	}

	return GameStatus(OnGoing)
}

func isTie(field [][]FieldValue) bool {
	for i := range field {
		for j := range field[0] {
			if field[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func isVictory(field [][]FieldValue, value FieldValue) bool {
	//Check horizontal
	horizontal_line := true
	for y := range field[0] {
		horizontal_line = true
		for x := range field {
			if field[x][y] != value {
				horizontal_line = false
			}
		}
		if horizontal_line {
			return true
		}
	}

	//Check vertical
	vertical_line := true
	for x := range field {
		vertical_line = true
		for y := range field[0] {
			if field[x][y] != value {
				vertical_line = false
			}
		}
		if vertical_line {
			return true
		}
	}

	//Check diagonal
	if field[0][2] == field[1][1] &&
		field[1][1] == field[2][0] &&
		field[2][0] == value {

		return true
	}

	if field[2][2] == field[1][1] &&
		field[1][1] == field[0][0] &&
		field[0][0] == value {

		return true
	}

	return false
}
