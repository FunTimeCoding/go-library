package assert

import (
	"fmt"
	"testing"
)

func TestGreater(t *testing.T) {
	Greater(t, 0, 1)

	// Unhappy less
	t1 := &testing.T{}
	Greater(t1, 1, 0)

	if !t1.Failed() {
		fmt.Println("unhappy less")
		t.Fail()
	}

	// Unhappy equal
	t2 := &testing.T{}
	Greater(t2, 1, 1)

	if !t2.Failed() {
		fmt.Println("unhappy equal")
		t.Fail()
	}
}
