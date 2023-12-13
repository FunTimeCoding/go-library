package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestIndexOfSkip(t *testing.T) {
	assert.Integer(
		t,
		1,
		IndexOfSkip(Bravo, []string{Alfa, Bravo, Charlie, Bravo}, 0),
	)
	assert.Integer(
		t,
		3,
		IndexOfSkip(Bravo, []string{Alfa, Bravo, Charlie, Bravo}, 1),
	)
}
