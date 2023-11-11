package expression

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestExpression(t *testing.T) {
	e := New([]string{strings.Alfa}, []string{strings.Delta})
	assert.True(t, e.Check([]string{strings.Alfa}))
	assert.True(t, e.Check([]string{strings.Alfa, strings.Bravo}))
	assert.False(t, e.Check([]string{strings.Alfa, strings.Delta}))
}
