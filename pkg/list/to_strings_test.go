package list

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

type Fixture string

func (f Fixture) String() string {
	return string(f)
}

func TestToStrings(t *testing.T) {
	assert.Strings(
		t,
		[]string{"Alfa", "Bravo"},
		ToStrings([]Fixture{upper.Alfa, upper.Bravo}),
	)
}
