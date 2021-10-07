package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"water-jug-riddle/riddle"
)

const (
	WelcomeMessage = "Welcome to the water jug riddle!"

	Instructions = `Instructions: You have an X-gallon and a Y-gallon jug that you can fill from a lake (You should assume the lake has unlimited amounts of water). 
By using only an X-gallon and a Y-gallon jug (no third jug), measure Z gallons of water.`

	EnterJugSize = "Please enter the size of Jug %s: "
	EnterDesired = "Please enter the desired volume to be reached: "
	Again        = "Do you want to play again? [y/n]: "
	GoodBye      = "Good bye!"

	Banner = "██╗    ██╗ █████╗ ████████╗███████╗██████╗          ██╗██╗   ██╗ ██████╗     ██████╗ ██╗██████╗ ██████╗ ██╗     ███████╗\n██║    ██║██╔══██╗╚══██╔══╝██╔════╝██╔══██╗         ██║██║   ██║██╔════╝     ██╔══██╗██║██╔══██╗██╔══██╗██║     ██╔════╝\n██║ █╗ ██║███████║   ██║   █████╗  ██████╔╝         ██║██║   ██║██║  ███╗    ██████╔╝██║██║  ██║██║  ██║██║     █████╗  \n██║███╗██║██╔══██║   ██║   ██╔══╝  ██╔══██╗    ██   ██║██║   ██║██║   ██║    ██╔══██╗██║██║  ██║██║  ██║██║     ██╔══╝  \n╚███╔███╔╝██║  ██║   ██║   ███████╗██║  ██║    ╚█████╔╝╚██████╔╝╚██████╔╝    ██║  ██║██║██████╔╝██████╔╝███████╗███████╗\n ╚══╝╚══╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝╚═╝  ╚═╝     ╚════╝  ╚═════╝  ╚═════╝     ╚═╝  ╚═╝╚═╝╚═════╝ ╚═════╝ ╚══════╝╚══════╝\n                                                                                                                        "
)

func PrintWelcomeMessage() {
	fmt.Println()
	fmt.Println(Banner)
	fmt.Println()
	fmt.Println(WelcomeMessage)
	fmt.Println()
	fmt.Println(Instructions)
}

func PrintGoodBye() {
	fmt.Println(GoodBye)
}

//ReadValues requests the input to the user until it gets 3 valid values. This is:
//1.- All values must be positive integers.
//2.- desired volume can't be higher than the biggest Jug
func ReadValues() (sizeX, sizeY, desired int64) {
	sizeX = readInteger(fmt.Sprintf(EnterJugSize, "X"))
	sizeY = readInteger(fmt.Sprintf(EnterJugSize, "Y"))
	desired = readInteger(EnterDesired)

	for true {
		if desired > sizeX && desired >sizeY {
			fmt.Println("Error: desired volume cannot be higher than the biggest Jug")
			desired = readInteger(EnterDesired)
		} else {
			break
		}

	}
	return sizeX, sizeY, desired
}

//readInteger reads the console until the user provides a positive integer.
func readInteger(msg string) int64 {
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Println()
		fmt.Print(msg)

		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		i, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil || i <= 0 {
			fmt.Println("Error: Only positive integers are allowed")
			continue
		}
		return i
	}
	return 0
}

//PlayAgain ask the user if he/she wants to play again.
func PlayAgain() bool {
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Println()
		fmt.Print(Again)

		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}
		input := strings.ToLower(scanner.Text())
		if len(input) != 1 || input != "y" && input != "n" {
			fmt.Println("Error: Only ['y', 'n', 'Y, 'N'] are allowed.")
			continue
		}
		return input == "y"
	}
	return false
}

//PrintSolution will print a solution if any, or the message "No Solution" otherwise.
func PrintSolution(step *riddle.Step){
	if step == nil {
		fmt.Println("No Solution")
	} else {
		fmt.Println()
		fmt.Println("The shortest way to achieve the desired volume is:")
		fmt.Println()
		doPrint(step)
	}
}

//doPrint recursively goes through the Steps of the solution and prints it in the corresponding order.
func doPrint(step *riddle.Step) {
	if step.Previous != nil {
		doPrint(step.Previous)
	}
	fmt.Println(step)
}
