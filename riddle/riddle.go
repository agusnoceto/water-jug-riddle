package riddle

func Solve(sizeX, sizeY, desired int64) *Step {
	if hasSolution(sizeX, sizeY, desired) {
		return doSolve(sizeX, sizeY, desired)
	}
	return nil
}

func doSolve(sizeX int64, sizeY int64, desired int64) *Step {
	queue := NewQueue()
	first := NewStep(sizeX, sizeY)
	queue.Push(*first)
	visited := make([]Step, 0)

	for !queue.IsEmpty() {
		current := queue.Pop()
		current.Execute()
		if contains(visited, current) {
			continue
		}

		if isSolution(current, desired) {
			return current
		}
		visited = append(visited, *current)
		nextSteps := current.NextSteps()
		queue.PushAll(nextSteps)
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

func isSolution(current *Step, requested int64) bool {
	return requested == current.JugX.amount || requested == current.JugY.amount
}

func hasSolution(sizeA, sizeB, requested int64) bool {
	gcd := resolveGCD(sizeA, sizeB)
	return requested%gcd == 0
}

func resolveGCD(numberA, numberB int64) int64 {
	if numberB == 0 {
		return numberA
	}
	return resolveGCD(numberB, numberA%numberB)
}
