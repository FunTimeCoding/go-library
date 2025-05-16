package scanner

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestScanner(t *testing.T) {
	assert.True(t, New() != nil)
}
