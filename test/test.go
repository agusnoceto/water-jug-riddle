package test

import (
	"fmt"
	"testing"
)

func AssertInts64(t *testing.T, label string, v1 int64, v2 int64) {
	if v1 != v2 {
		t.Logf(fmt.Sprintf("%s: %v != %v", label, v1, v2))
		t.Fail()
	}
}

func AssertInts(t *testing.T, label string, v1 int, v2 int) {
	if v1 != v2 {
		t.Logf(fmt.Sprintf("%s: %v != %v", label, v1, v2))
		t.Fail()
	}
}

func AssertBool(t *testing.T, label string, v1 bool) {
	if !v1 {
		t.Logf(fmt.Sprintf("%s: %v", label, v1))
		t.Fail()
	}
}

func AssertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Logf("Func did not panic")
			t.Fail()
		}
	}()
	f()
}
