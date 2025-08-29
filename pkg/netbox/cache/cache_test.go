package cache

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestCache(t *testing.T) {
	assert.NotNil(t, New())
}
