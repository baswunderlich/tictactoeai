package main

import (
	"fmt"
	"tictactoeai/ai"
	"tictactoeai/tictactoe"
)

func main() {

	//Here you can define the AIs used
	brainCross := ai.RandomAI{}
	brainCircle := ai.RandomAI{}

	gameRuns := 100
	score := 0

	for range gameRuns {
		game := tictactoe.CreateGame()
		game.PrintField()

		//Game loop
		for {
			var nextMove ai.Coordinate

			if game.CurrentPlayer == tictactoe.Cross {
				fmt.Println("Its the turn of player Cross")
				nextMove = brainCross.GetMove(game)
			} else {
				fmt.Println("Its the turn of player Circle")
				nextMove = brainCircle.GetMove(game)
			}

			x := nextMove.X
			y := nextMove.Y

			if game.CurrentPlayer == tictactoe.Cross {
				game.PlaceCross(x, y)
			} else {
				game.PlaceCircle(x, y)
			}
			game.PrintField()
			status := game.GetStatus()

			if status != tictactoe.GameStatus(tictactoe.OnGoing) {
				score += int(status)
				fmt.Printf("Game finished: %s\n", status.String())
				fmt.Println("--------------------------------------")
				break
			}
		}
	}

	fmt.Printf("Final Score: %d", score)
}
