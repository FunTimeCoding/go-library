package text

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestOptimizeWhitespace(t *testing.T) {
	assert.String(
		t,
		"",
		OptimizeWhitespace("", 1),
	)
	assert.String(
		t,
		"A\nB\n",
		OptimizeWhitespace("A\nB\n", 1),
	)
	assert.String(
		t,
		"A\n\nB\n",
		OptimizeWhitespace("A\n\nB\n", 1),
	)
	assert.String(
		t,
		"A\n\nB\n",
		OptimizeWhitespace("A\n\n\nB\n", 1),
	)
	assert.String(
		t,
		"A\n\nB\n",
		OptimizeWhitespace("A\n \n \nB\n\n", 1),
	)
}
