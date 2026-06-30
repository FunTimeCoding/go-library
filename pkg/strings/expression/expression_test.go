package expression

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestExpression(t *testing.T) {
	e := New([]string{upper.Alfa}, []string{upper.Delta})
	assert.True(t, e.Check([]string{upper.Alfa}))
	assert.True(t, e.Check([]string{upper.Alfa, upper.Bravo}))
	assert.False(t, e.Check([]string{upper.Alfa, upper.Delta}))
}
