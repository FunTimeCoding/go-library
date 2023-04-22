package time

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
	"time"
)

func TestOlderThan(t *testing.T) {
	now := time.Now()
	assert.Boolean(t, false, OlderThan(now, 1))

	fiveMinutesAgo := now.Add(-5 * time.Minute)
	assert.Boolean(t, true, OlderThan(fiveMinutesAgo, 1))
}
