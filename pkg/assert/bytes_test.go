package assert

import "testing"

func TestBytes(t *testing.T) {
	Bytes(t, []byte("123"), []byte("123"))
}
