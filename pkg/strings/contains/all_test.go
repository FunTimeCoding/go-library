package contains

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestAll(t *testing.T) {
	assert.True(t, All([]string{strings.Alfa}, []string{strings.Alfa}))
	assert.False(t, All([]string{strings.Alfa}, []string{strings.Bravo}))

	assert.True(
		t,
		All(
			[]string{strings.Alfa, strings.Bravo, strings.Charlie},
			[]string{strings.Alfa, strings.Bravo},
		),
	)
	assert.False(
		t,
		All(
			[]string{strings.Alfa, strings.Bravo, strings.Charlie},
			[]string{strings.Alfa, strings.Delta},
		),
	)
}
