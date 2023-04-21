package assert

import "testing"

func TestRound(t *testing.T) {
	Round(t, 1.1, 1.14, 1)
	Round(t, 1.15, 1.154, 2)
}
