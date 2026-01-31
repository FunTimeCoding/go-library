package markdown

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestLink(t *testing.T) {
	assert.String(
		t,
		"[Alfa](Bravo)",
		Link(strings.Alfa, strings.Bravo),
	)
}
