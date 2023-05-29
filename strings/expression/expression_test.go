package expression

import (
	"github.com/funtimecoding/go-library/assert"
	"github.com/funtimecoding/go-library/strings"
	"testing"
)

func TestExpression(t *testing.T) {
	e := New([]string{strings.Alfa}, []string{strings.Delta})
	assert.True(t, e.Evaluate([]string{strings.Alfa}))
	assert.True(t, e.Evaluate([]string{strings.Alfa, strings.Bravo}))
	assert.False(t, e.Evaluate([]string{strings.Alfa, strings.Delta}))
}
