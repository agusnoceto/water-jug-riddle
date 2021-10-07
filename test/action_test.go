package test

import (
	"testing"
	"water-jug-riddle/riddle"
)

func TestActions(t *testing.T) {
	actions := riddle.Actions()
	AssertInts(t, "Expected actions to be 6", len(actions), 6)
}
