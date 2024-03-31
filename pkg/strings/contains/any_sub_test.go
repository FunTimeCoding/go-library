package contains

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestSubAny(t *testing.T) {
	assert.True(t, AnySub([]string{strings.Alfa}, []string{"Al"}))
	assert.False(t, AnySub([]string{strings.Alfa}, []string{"Ga"}))
	assert.False(t, AnySub([]string{strings.Alfa}, []string{}))
	assert.False(t, AnySub([]string{}, []string{}))
	assert.False(t, AnySub([]string{}, []string{"Al"}))
	assert.True(
		t,
		AnySub(
			[]string{strings.Alfa, strings.Bravo, strings.Charlie},
			[]string{"Br"},
		),
	)
	assert.False(
		t,
		AnySub(
			[]string{strings.Alfa, strings.Bravo, strings.Charlie},
			[]string{"De"},
		),
	)
}
