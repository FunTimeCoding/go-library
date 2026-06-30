package contains

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestSubAny(t *testing.T) {
	assert.True(t, AnySub([]string{"Al"}, []string{upper.Alfa}))
	assert.False(t, AnySub([]string{"Ga"}, []string{upper.Alfa}))
	assert.False(t, AnySub([]string{}, []string{upper.Alfa}))
	assert.False(t, AnySub([]string{}, []string{}))
	assert.False(t, AnySub([]string{"Al"}, []string{}))
	assert.True(
		t,
		AnySub(
			[]string{"Br"},
			[]string{upper.Alfa, upper.Bravo, upper.Charlie},
		),
	)
	assert.False(
		t,
		AnySub(
			[]string{"De"},
			[]string{upper.Alfa, upper.Bravo, upper.Charlie},
		),
	)
}
