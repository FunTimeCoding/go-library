package expression

import (
	assert2 "github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestExpression(t *testing.T) {
	e := New([]string{strings.Alfa}, []string{strings.Delta})
	assert2.True(t, e.Check([]string{strings.Alfa}))
	assert2.True(t, e.Check([]string{strings.Alfa, strings.Bravo}))
	assert2.False(t, e.Check([]string{strings.Alfa, strings.Delta}))
}
