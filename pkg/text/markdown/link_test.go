package markdown

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestLink(t *testing.T) {
	assert.String(
		t,
		"[Alfa](Bravo)",
		Link(upper.Alfa, upper.Bravo),
	)
}
