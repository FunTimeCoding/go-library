package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestStripUntilDirectory(t *testing.T) {
	assert.String(
		t,
		"/test",
		StripUntilDirectory("/some/long/path/test", "test"),
	)
}
