package test

import (
	"testing"
	"water-jug-riddle/riddle"
)

func TestSolve(t *testing.T) {
	var desired int64 = 4
	solution := riddle.Solve(5, 3, desired)
	AssertBool(t, "Expected solution to be found", solution.JugX.Amount() == desired || solution.JugY.Amount() == desired)

	solution = riddle.Solve(4, 2, 1)
	AssertBool(t, "Expected No solution", solution == nil)

}
