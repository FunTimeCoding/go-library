package contains

import (
	assert2 "github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestAny(t *testing.T) {
	assert2.True(t, Any([]string{strings.Alfa}, []string{strings.Alfa}))
	assert2.False(t, Any([]string{strings.Alfa}, []string{strings.Bravo}))

	assert2.True(
		t,
		Any(
			[]string{strings.Alfa, strings.Bravo, strings.Charlie},
			[]string{strings.Alfa},
		),
	)
	assert2.False(
		t,
		Any(
			[]string{strings.Alfa, strings.Bravo, strings.Charlie},
			[]string{strings.Delta},
		),
	)
}
