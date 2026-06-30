package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestIndexOf(t *testing.T) {
	assert.Integer(
		t,
		1,
		IndexOf(
			upper.Bravo,
			[]string{upper.Alfa, upper.Bravo, upper.Charlie},
		),
	)
}
