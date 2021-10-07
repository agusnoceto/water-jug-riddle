package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"water-jug-riddle/riddle"
)

const (
	Banner         = "██╗    ██╗ █████╗ ████████╗███████╗██████╗          ██╗██╗   ██╗ ██████╗     ██████╗ ██╗██████╗ ██████╗ ██╗     ███████╗\n██║    ██║██╔══██╗╚══██╔══╝██╔════╝██╔══██╗         ██║██║   ██║██╔════╝     ██╔══██╗██║██╔══██╗██╔══██╗██║     ██╔════╝\n██║ █╗ ██║███████║   ██║   █████╗  ██████╔╝         ██║██║   ██║██║  ███╗    ██████╔╝██║██║  ██║██║  ██║██║     █████╗  \n██║███╗██║██╔══██║   ██║   ██╔══╝  ██╔══██╗    ██   ██║██║   ██║██║   ██║    ██╔══██╗██║██║  ██║██║  ██║██║     ██╔══╝  \n╚███╔███╔╝██║  ██║   ██║   ███████╗██║  ██║    ╚█████╔╝╚██████╔╝╚██████╔╝    ██║  ██║██║██████╔╝██████╔╝███████╗███████╗\n ╚══╝╚══╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝╚═╝  ╚═╝     ╚════╝  ╚═════╝  ╚═════╝     ╚═╝  ╚═╝╚═╝╚═════╝ ╚═════╝ ╚══════╝╚══════╝\n                                                                                                                        "
	WelcomeMessage = "Welcome to the water jug riddle!"

	Instructions = `Instructions: You have an X-gallon and a Y-gallon jug that you can fill from a lake (You should assume the lake has unlimited amounts of water). 
By using only an X-gallon and a Y-gallon jug (no third jug), measure Z gallons of water.`

	EnterJugSize     = "Please enter the size of Jug %s: "
	EnterDesired     = "Please enter the desired volume to be reached: "
	Again            = "Do you want to play again? [y/n]: "
	GoodBye          = "Good bye!"
	NoSolution       = "No Solution"
	SolutionHeader   = "The shortest way to achieve the desired volume is:"
	NumberOfSteps    = "\nNumber of Steps: %d\n"
	SizeThreshold    = 50000
	ThresholdWarning = "Some of the values entered are large numbers. This might take some time. Find solution anyway? [y/n]: "
	ExecutionTime    = "Process took %s to finish\n"
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
//2.- desired volume can't be higher than the biggest Jug.
//3.- If large numbers are entered, a warning message is displayed letting the user know it might take time.
func ReadValues() (sizeX, sizeY, desired int64) {
	for true {
		sizeX = readInteger(fmt.Sprintf(EnterJugSize, "X"))
		sizeY = readInteger(fmt.Sprintf(EnterJugSize, "Y"))
		desired = readInteger(EnterDesired)

		//desired <= max(sizeX, sizeY)
		desired = readValidDesired(sizeX, sizeY, desired)

		//if large numbers, ask the user if he/she wants to continue
		if sizeX > SizeThreshold || sizeY > SizeThreshold || desired > SizeThreshold {
			if readYes(ThresholdWarning) {
				return sizeX, sizeY, desired
			}
		} else {
			return sizeX, sizeY, desired
		}
	}
	return sizeX, sizeY, desired
}

//readValidDesired will keep asking the user for a number less <= max(sizeX, sizeY)
func readValidDesired(sizeX int64, sizeY int64, desired int64) int64 {
	for true {
		if desired <= sizeX || desired <= sizeY {
			break
		}
		fmt.Println("Error: desired volume cannot be larger than the biggest Jug")
		desired = readInteger(EnterDesired)
	}
	return desired
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


//readYes returns true if the user enters 'y' or 'Y'.
//false if he/she enters 'n', 'N'. No others characters allowed.
func readYes(msg string) bool {
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Println()
		fmt.Print(msg)

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

func PlayAgain() bool {
	return readYes(Again)
}

//PrintSolution will print a solution if any, or the message "No Solution" otherwise.
func PrintSolution(step *riddle.Step, elapsed time.Duration) {
	if step == nil {
		fmt.Println(NoSolution)
	} else {
		fmt.Println()
		fmt.Println(SolutionHeader)
		fmt.Println()
		steps := doPrint(step)
		fmt.Printf(NumberOfSteps, steps)
		fmt.Printf(ExecutionTime, elapsed)
	}
}

//doPrint recursively goes through the Steps of the solution and prints it in the corresponding order.
//it returns the number of Steps.
func doPrint(step *riddle.Step) int64 {
	var steps int64
	if step.Previous != nil {
		steps = 1 + doPrint(step.Previous)
	}
	fmt.Println(step)
	return steps
}
