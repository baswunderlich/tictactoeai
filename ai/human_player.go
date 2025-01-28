package ai

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"tictactoeai/tictactoe"
)

type HumanAI struct{}

func (r *HumanAI) GetMove(game tictactoe.Game) Coordinate {
	for {
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatal(err)
		}
		coords := strings.Split(input, ",")
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			continue
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			continue
		}
		return Coordinate{X: x, Y: y}
	}
}
