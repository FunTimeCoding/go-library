package assert

import "testing"

func TestFloat(t *testing.T) {
	Float(t, 1.1, 1.1)
	Float(t, 1.15, 1.15)
}
