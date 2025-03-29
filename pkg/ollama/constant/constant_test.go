package constant

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.Integer(t, 11434, Port)
	assert.String(t, "llama3.2", Llama32)
	assert.String(t, "llama3.2:1b", Llama321b)
}
