package issue

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestHasKey(t *testing.T) {
	assert.True(
		t,
		HasKey("Message with key ABC-123."),
	)
}
