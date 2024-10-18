package output

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestSettings(t *testing.T) {
	assert.Any(t, &Settings{Format: "text"}, New())
	assert.Any(
		t,
		&Settings{Format: "text", Debug: true},
		New(WithDebug()),
	)
	assert.Any(
		t,
		&Settings{Format: "markdown"},
		New(WithFormat(Markdown)),
	)
	assert.Any(
		t,
		&Settings{Format: "notation"},
		New(WithFormat(Notation)),
	)
}
