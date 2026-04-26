package assert

import "testing"

var StubCount int

func Stub(t *testing.T) {
	t.Helper()
	StubCount++
}
