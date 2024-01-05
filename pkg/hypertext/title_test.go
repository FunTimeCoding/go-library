package hypertext

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestTitle(t *testing.T) {
	assert.String(
		t,
		"Test Title",
		Title(Document(fixtureFile("test.html"))),
	)
}
