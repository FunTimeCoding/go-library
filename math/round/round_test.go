package round

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestRound(t *testing.T) {
	assert.Equal(
		t,
		1.0,
		Round(1.01, 1),
	)
}
