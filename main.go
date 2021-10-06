package main

import (
	"water-jug-riddle/riddle"
	"water-jug-riddle/ui"
)

func main() {

	ui.PrintWelcomeMessage()
	play := true

	for  play {

		sizeX, sizeY, desired := ui.ReadValues()
		solution := riddle.Solve(sizeX, sizeY, desired)
		ui.PrintSolution(solution)
		play = ui.PlayAgain()
	}
	ui.PrintGoodBye()
}
