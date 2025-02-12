package contains

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestSubAny(t *testing.T) {
	assert.True(t, AnySub([]string{"Al"}, []string{strings.Alfa}))
	assert.False(t, AnySub([]string{"Ga"}, []string{strings.Alfa}))
	assert.False(t, AnySub([]string{}, []string{strings.Alfa}))
	assert.False(t, AnySub([]string{}, []string{}))
	assert.False(t, AnySub([]string{"Al"}, []string{}))
	assert.True(
		t,
		AnySub(
			[]string{"Br"},
			[]string{strings.Alfa, strings.Bravo, strings.Charlie},
		),
	)
	assert.False(
		t,
		AnySub(
			[]string{"De"},
			[]string{strings.Alfa, strings.Bravo, strings.Charlie},
		),
	)
}
