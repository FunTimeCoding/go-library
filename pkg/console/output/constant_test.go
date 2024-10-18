package output

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "notation", Notation)
	assert.String(t, "markdown", Markdown)
	assert.True(t, Default != nil)
	assert.True(t, Debug != nil)
}
