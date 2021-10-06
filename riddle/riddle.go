package riddle

//Solve will determine if there's a solution for the given input. If there is one,
//it will return the last Step of that solution. if not it will return nil.
func Solve(sizeX, sizeY, desired int64) *Step {
	if hasSolution(sizeX, sizeY, desired) {
		return doSolve(sizeX, sizeY, desired)
	}
	return nil
}

//doSolve finds a solution using the following heuristic:
//Create an N-ary tree in which each node (Step) of the tree represents a state of 2 Jugs in a given time.
//The Root node will contain a Step representing the initial State: i.e. Two Empty Jugs of the sizes entered by the user.
//Each following Level in the tree will contain all valid moves (Action) of water. This will be the Children of Root.
//
//So the second level will contain the 2 possible first moves (Action). This is:
//One Step with Action = FillX (fill Jug X from the lake).
//One Step with Action = FillY (fill Jug Y from the lake).
//
//Third Level will have the valid moves (Action) of Both of those Steps, and so on.
//
//Thus, each Level represents the amount of moves needed to reach the Steps of said Level.
//
//By traversing that tree in BFS, the first solution found is the shortest one,
//since each Level is fully traversed before moving into the next one.
//
//It iterates the N-ary tree in BFS by using a Queue starting with the Root Step.
//If Root is not the Solution, then it will remove it from the queue and add its Children.
//The process will repeat until it finds a Solution or the Queue is empty.
func doSolve(sizeX int64, sizeY int64, desired int64) *Step {
	stepsQueue := NewQueue()
	first := NewStep(sizeX, sizeY)
	stepsQueue.Push(*first)
	visited := make([]Step, 0)

	for !stepsQueue.IsEmpty() {
		current := stepsQueue.Pop()
		current.Execute() //executes the Step action
		if contains(visited, current) {
			continue
		}

		if isSolution(current, desired) {
			return current
		}

		visited = append(visited, *current)
		stepsQueue.PushAll(current.NextSteps())
	}
	return nil
}

func contains(steps []Step, step *Step) bool {
	for _, s := range steps {
		if s.Equals(step) {
			return true
		}
	}
	return false
}

//isSolution checks whether the desired volume is in any of the 2 Jugs of this Step.
func isSolution(current *Step, requested int64) bool {
	return requested == current.JugX.amount || requested == current.JugY.amount
}

//checks if the given input has solutions by checking if the desired amount is divisible by
//the Greatest Common Divisor of the Jug sizes.
func hasSolution(sizeA, sizeB, requested int64) bool {
	gcd := resolveGCD(sizeA, sizeB)
	return requested%gcd == 0
}

//resolveGCD calculates the Greates Common Divisor of the 2 arguments.
func resolveGCD(numberA, numberB int64) int64 {
	if numberB == 0 {
		return numberA
	}
	return resolveGCD(numberB, numberA%numberB)
}
