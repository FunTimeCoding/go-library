package strings

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestContainsAny(t *testing.T) {
	assert.True(t, ContainsAny([]string{Alfa}, []string{Alfa}))
	assert.False(t, ContainsAny([]string{Alfa}, []string{Bravo}))

	assert.True(
		t,
		ContainsAny(
			[]string{Alfa, Bravo, Charlie},
			[]string{Alfa},
		),
	)
	assert.False(
		t,
		ContainsAny(
			[]string{Alfa, Bravo, Charlie},
			[]string{Delta},
		),
	)
}
