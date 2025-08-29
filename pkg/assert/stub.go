package assert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var StubCount int

func Stub(t *testing.T) {
	t.Helper()
	StubCount++
	assert.True(t, true)
}
