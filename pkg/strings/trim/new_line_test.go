package trim

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestNewLine(t *testing.T) {
	assert.String(t, "hello", NewLine("hello\n"))
}
