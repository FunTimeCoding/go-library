package assert

import "testing"

func TestIntegers(t *testing.T) {
	Integers(t, []int{0, 1}, []int{0, 1})
}
