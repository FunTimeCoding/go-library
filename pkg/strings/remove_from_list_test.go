package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestRemoveFromList(t *testing.T) {
	assert.Any(
		t,
		[]string{Bravo, Charlie},
		RemoveFromList(
			[]string{Alfa, Alfa, Bravo, Charlie},
			[]string{Alfa},
		),
	)
	assert.Any(
		t,
		[]string{Charlie},
		RemoveFromList(
			[]string{Alfa, Alfa, Bravo, Charlie},
			[]string{Alfa, Bravo},
		),
	)
}
