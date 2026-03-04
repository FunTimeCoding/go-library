package assert

import (
	"gotest.tools/v3/assert"
	"testing"
	"time"
)

func Duration(
	t *testing.T,
	expect time.Duration,
	actual time.Duration,
) {
	t.Helper()
	assert.Equal(t, actual, expect)
}
