package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestReverse(t *testing.T) {
	reversed := []string{upper.Charlie, upper.Bravo, upper.Alfa}
	Reverse(reversed)
	assert.Any(t, []string{upper.Alfa, upper.Bravo, upper.Charlie}, reversed)
}
