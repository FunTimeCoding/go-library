package assert

import (
	"fmt"
	"testing"
)

func TestLessEqual(t *testing.T) {
	LessEqual(t, 1, 0)
	LessEqual(t, 1, 1)

	// Unhappy more
	t1 := &testing.T{}
	LessEqual(t1, 0, 1)

	if !t1.Failed() {
		fmt.Println("unhappy more")
		t.Fail()
	}
}
