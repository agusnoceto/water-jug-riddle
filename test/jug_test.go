package test

import (
	"fmt"
	"testing"
	"water-jug-riddle/riddle"
)

func TestIsEmpty(t *testing.T) {
	jug := riddle.NewJug(5, 0)
	AssertBool(t, "Expected: Jug to be empty", jug.IsEmpty())

	jug = riddle.NewJug(5, 1)
	AssertBool(t, "Expected: Jug not to be empty", !jug.IsEmpty())
}

func TestIsFull(t *testing.T) {
	jug := riddle.NewJug(5, 5)
	AssertBool(t, "Expected: Jug to be full", jug.IsFull())

	jug = riddle.NewJug(5, 1)
	AssertBool(t, "Expected: Jug not to be full", !jug.IsFull())
}

func TestRemaining(t *testing.T) {
	var size int64 = 5
	var amount int64 = 3
	remaining := size - amount
	jug := riddle.NewJug(size, amount)
	AssertInts64(t, fmt.Sprintf("Expected: remaining to be %d", remaining), jug.Remaining(), remaining)
}

func TestIsFill(t *testing.T) {
	var size int64 = 5
	jug := riddle.NewJug(size, 0)
	jug.Fill()
	AssertBool(t, "Expected: jug to by full", jug.IsFull())
}

func TestEmpty(t *testing.T) {
	var size int64 = 5
	jug := riddle.NewJug(size, size)
	jug.Empty()
	AssertBool(t, "Expected: jug to by empty", jug.IsEmpty())

	jug = riddle.NewJug(size, 0)
	AssertPanic(t, jug.Empty)
}

func TestPour(t *testing.T) {
	var size int64 = 5
	var amount int64 = 3
	var volume int64 = 4

	jug := riddle.NewJug(size, amount)
	remainder := jug.Pour(volume)
	AssertInts64(t, fmt.Sprintf("Expected remainder to be %d", volume-(size-amount)), remainder, volume-(size-amount))
	AssertBool(t, "Expected Jug to be full", jug.IsFull())

	jug = riddle.NewJug(12, 3)
	remainder = jug.Pour(volume)
	AssertInts64(t, fmt.Sprintf("Expected remainder to be 0"), 0, remainder)
	AssertBool(t, "Expected: jug not to be full", !jug.IsFull())

	defer func() {
		if r := recover(); r == nil {
			t.Logf("Func did not panic")
			t.Fail()
		}
	}()
	jug.Pour(0)
}

func TestSet(t *testing.T) {
	var size int64 = 5
	var amount int64 = 4

	jug := riddle.NewJug(size, 0)
	jug.Set(amount)
	AssertInts64(t, fmt.Sprintf("Expected: remaining to be %d", size-amount), jug.Remaining(), size-amount)

	defer func() {
		if r := recover(); r == nil {
			t.Logf("Func did not panic")
			t.Fail()
		}
	}()
	jug.Set(-1)
}

func TestJugEquals(t *testing.T) {
	var size int64 = 5
	var amount int64 = 4
	jug := riddle.NewJug(size, amount)
	other := riddle.NewJug(size, amount)

	AssertBool(t, "Expected jugs to be equal in size and amount", jug.Equals(*other))
	other.Empty()
	AssertBool(t, "Expected jugs to be different in amount", !jug.Equals(*other))
}

func TestClone(t *testing.T) {
	jug := riddle.NewJug(8, 6)
	clone := jug.Clone()
	AssertBool(t, "Expected structs to be different", clone != jug)
	AssertInts64(t, "Expected remaining to be equals", clone.Remaining(), jug.Remaining())
}
