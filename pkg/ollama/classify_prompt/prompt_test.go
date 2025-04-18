package classify_prompt

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestPrompt(t *testing.T) {
	assert.True(t, New() != nil)
}
