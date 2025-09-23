package assert

import "testing"

func TestNil(t *testing.T) {
	Nil(t, nil)
}
