package check_entry

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestEntry(t *testing.T) {
	assert.True(t, New(Warning, "test") != nil)
}
