package contains

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestAll(t *testing.T) {
	assert.True(t, All([]string{upper.Alfa}, []string{upper.Alfa}))
	assert.False(t, All([]string{upper.Alfa}, []string{upper.Bravo}))
	assert.True(
		t,
		All(
			[]string{upper.Alfa, upper.Bravo, upper.Charlie},
			[]string{upper.Alfa, upper.Bravo},
		),
	)
	assert.False(
		t,
		All(
			[]string{upper.Alfa, upper.Bravo, upper.Charlie},
			[]string{upper.Alfa, upper.Delta},
		),
	)
}
