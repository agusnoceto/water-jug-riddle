package test

import (
	"testing"
	"water-jug-riddle/riddle"
)

func TestPush(t *testing.T) {
	queue := riddle.NewQueue()
	AssertBool(t, "Expected Size to be 0", queue.Size() == 0)

	step := riddle.NewStep(3, 4)
	queue.Push(*step)
	AssertBool(t, "Expected Size to be 1", queue.Size() == 1)
}

func TestPushAll(t *testing.T) {

	queue := riddle.NewQueue()
	AssertBool(t, "Expected Size to be 0", queue.Size() == 0)

	step1 := riddle.NewStep(3, 4)
	step2 := riddle.NewStep(3, 5)
	queue.PushAll([]riddle.Step{*step1, *step2})
	AssertBool(t, "Expected Size to be 2", queue.Size() == 2)
}

func TestPop(t *testing.T) {

	queue := riddle.NewQueue()
	step := riddle.NewStep(3, 4)
	queue.Push(*step)
	AssertBool(t, "Expected Size to be 1", queue.Size() == 1)
	pop := queue.Pop()
	AssertBool(t, "Expected Size to be 0", queue.Size() == 0)
	AssertBool(t, "Expected step to be equals", step.Equals(pop))
}

func TestSize(t *testing.T) {
	queue := riddle.NewQueue()
	AssertBool(t, "Expected Size to be 0", queue.Size() == 0)
	step := riddle.NewStep(3, 4)
	queue.Push(*step)
	AssertBool(t, "Expected Size to be 1", queue.Size() == 1)
}

func TestIsQueueEmpty(t *testing.T) {
	queue := riddle.NewQueue()
	AssertBool(t, "Expected IsEmpty() to be true", queue.IsEmpty())
	step := riddle.NewStep(3, 4)
	queue.Push(*step)
	AssertBool(t, "Expected IsEmpty() to be false", !queue.IsEmpty() )
}