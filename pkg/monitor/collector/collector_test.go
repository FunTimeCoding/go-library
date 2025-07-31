package collector

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestCollector(t *testing.T) {
	assert.True(t, New("", "", "", 0, nil) != nil)
}
