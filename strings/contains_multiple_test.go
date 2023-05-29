package strings

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestContainsMultiple(t *testing.T) {
	assert.True(t, ContainsMultiple([]string{Alfa}, []string{Alfa}))
	assert.False(t, ContainsMultiple([]string{Alfa}, []string{Bravo}))

	assert.True(
		t,
		ContainsMultiple(
			[]string{Alfa, Bravo},
			[]string{Alfa, Bravo, Charlie},
		),
	)
	assert.False(
		t,
		ContainsMultiple(
			[]string{Alfa, Delta},
			[]string{Alfa, Bravo, Charlie},
		),
	)
}
