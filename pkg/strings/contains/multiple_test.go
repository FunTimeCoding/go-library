package contains

import (
	assert2 "github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestMultiple(t *testing.T) {
	assert2.True(t, Multiple([]string{strings.Alfa}, []string{strings.Alfa}))
	assert2.False(t, Multiple([]string{strings.Alfa}, []string{strings.Bravo}))

	assert2.True(
		t,
		Multiple(
			[]string{strings.Alfa, strings.Bravo},
			[]string{strings.Alfa, strings.Bravo, strings.Charlie},
		),
	)
	assert2.False(
		t,
		Multiple(
			[]string{strings.Alfa, strings.Delta},
			[]string{strings.Alfa, strings.Bravo, strings.Charlie},
		),
	)
}
