package main

import (
	"time"
	"water-jug-riddle/riddle"
	"water-jug-riddle/ui"
)

func main() {
	ui.PrintWelcomeMessage()

	play := true
	for play {
		sizeX, sizeY, desired := ui.ReadValues()

		start := time.Now()
		solution := riddle.Solve(sizeX, sizeY, desired)

		ui.PrintSolution(solution, time.Since(start))

		play = ui.PlayAgain()
	}
	ui.PrintGoodBye()
}
