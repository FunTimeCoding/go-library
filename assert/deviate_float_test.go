package assert

import "testing"

func TestDeviate(t *testing.T) {
	Deviate(t, 1, 1.1, 0.1)
	Deviate(t, -1, -1.1, 0.1)
}
