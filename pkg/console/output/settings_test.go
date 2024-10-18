package output

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestSettings(t *testing.T) {
	assert.True(t, Ensure(nil) != nil)
	assert.True(t, Ensure(New()) != nil)
	assert.Any(t, &Settings{Format: "text"}, FromOptions())
	assert.Any(
		t,
		&Settings{Format: "text", ShowDebug: true},
		FromOptions(WithDebug()),
	)
}
