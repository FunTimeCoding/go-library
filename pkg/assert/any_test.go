package assert

import "testing"

type Fixture struct {
	Value string
}

func TestAny(t *testing.T) {
	Any(t, &Fixture{Value: "a"}, &Fixture{Value: "a"})
}
