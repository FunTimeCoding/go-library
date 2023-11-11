package contains

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestMultiple(t *testing.T) {
	assert.True(t, Multiple([]string{strings.Alfa}, []string{strings.Alfa}))
	assert.False(t, Multiple([]string{strings.Alfa}, []string{strings.Bravo}))

	assert.True(
		t,
		Multiple(
			[]string{strings.Alfa, strings.Bravo},
			[]string{strings.Alfa, strings.Bravo, strings.Charlie},
		),
	)
	assert.False(
		t,
		Multiple(
			[]string{strings.Alfa, strings.Delta},
			[]string{strings.Alfa, strings.Bravo, strings.Charlie},
		),
	)
}
