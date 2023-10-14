package assert

import (
	"fmt"
	"testing"
)

func TestGreaterEqual(t *testing.T) {
	GreaterEqual(t, 0, 1)
	GreaterEqual(t, 1, 1)

	// Unhappy less
	t1 := &testing.T{}
	GreaterEqual(t1, 1, 0)

	if !t1.Failed() {
		fmt.Println("unhappy less")
		t.Fail()
	}
}
