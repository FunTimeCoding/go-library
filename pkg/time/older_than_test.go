package time

import (
	assert2 "github.com/funtimecoding/go-library/pkg/assert"
	"testing"
	"time"
)

func TestOlderThan(t *testing.T) {
	now := time.Now()
	assert2.False(t, OlderThan(now, 1))

	fiveMinutesAgo := now.Add(-5 * time.Minute)
	assert2.True(t, OlderThan(fiveMinutesAgo, 1))
}
