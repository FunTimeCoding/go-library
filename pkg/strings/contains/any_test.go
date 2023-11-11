package contains

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestAny(t *testing.T) {
	assert.True(t, Any([]string{strings.Alfa}, []string{strings.Alfa}))
	assert.False(t, Any([]string{strings.Alfa}, []string{strings.Bravo}))

	assert.True(
		t,
		Any(
			[]string{strings.Alfa, strings.Bravo, strings.Charlie},
			[]string{strings.Alfa},
		),
	)
	assert.False(
		t,
		Any(
			[]string{strings.Alfa, strings.Bravo, strings.Charlie},
			[]string{strings.Delta},
		),
	)
}
