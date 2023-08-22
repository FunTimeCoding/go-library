package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestIndexOf(t *testing.T) {
	assert.Integer(
		t,
		1,
		IndexOf(
			Bravo,
			[]string{Alfa, Bravo, Charlie},
		),
	)
}
