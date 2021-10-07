package test

import (
	"testing"
	"water-jug-riddle/riddle"
)

func TestExecute(t *testing.T) {
	step := riddle.NewStep(9, 2)
	step.Action = riddle.FillX

	AssertBool(t, "Expected jug X to be empty", step.JugX.IsEmpty())
	step.Execute()
	AssertBool(t, "Expected jug X to be full", step.JugX.IsFull())
}

func TestNextSteps(t *testing.T) {
	step := riddle.NewStep(9, 2)
	nextSteps := step.NextSteps()
	AssertBool(t, "Expected next steps length to be 2", len(nextSteps) == 2)
	var fillX, fillY bool
	for _, next := range nextSteps {
		if next.Action == riddle.FillX {
			fillX = true
		}
		if next.Action == riddle.FillY {
			fillY = true
		}
	}
	AssertBool(t, "Expected FillX to be one of the next steps", fillX)
	AssertBool(t, "Expected FillY to be one of the next steps", fillY)
}

func TestStepEquals(t *testing.T) {
	var size int64 = 8
	var amount int64 = 6
	step := riddle.NewStep(size, amount)
	equal := riddle.NewStep(size, amount)
	AssertBool(t, "Expected: Steps should be equals", step.Equals(equal))

	different := riddle.NewStep(size, amount-1)
	AssertBool(t, "Expected: Steps should be different", !step.Equals(different))

}
