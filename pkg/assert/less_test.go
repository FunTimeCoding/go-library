package assert

import (
	"fmt"
	"testing"
)

func TestLess(t *testing.T) {
	Less(t, 1, 0)

	// Unhappy more
	t1 := &testing.T{}
	Less(t1, 0, 1)

	if !t1.Failed() {
		fmt.Println("unhappy more")
		t.Fail()
	}

	// Unhappy equal
	t2 := &testing.T{}
	Less(t2, 1, 1)

	if !t2.Failed() {
		fmt.Println("unhappy equal")
		t.Fail()
	}
}
