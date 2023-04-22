package strings

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	reversed := []string{Charlie, Bravo, Alfa}
	Reverse(reversed)
	assert.Any(t, []string{Alfa, Bravo, Charlie}, reversed)
}
