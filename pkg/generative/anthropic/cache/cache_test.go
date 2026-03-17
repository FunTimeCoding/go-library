package cache

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestCache(t *testing.T) {
	assert.Any(t, Mode(1), FiveMinute)
}
