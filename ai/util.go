package ai

import (
	"tictactoeai/tictactoe"
)

type Coordinate struct {
	X int
	Y int
}

func GetPossibleMoves(game tictactoe.Game) []Coordinate {
	freeCoordinates := make([]Coordinate, 1)

	for x := range game.Field {
		for y := range game.Field[0] {
			if game.Field[x][y] == 0 {
				freeCoordinates = append(freeCoordinates,
					Coordinate{X: x, Y: y})
			}
		}
	}
	return freeCoordinates
}
