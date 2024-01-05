package hypertext

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestTables(t *testing.T) {
	assert.Any(
		t,
		[]string{"Example TD"},
		Tables(Document(fixtureFile("test.html"))),
	)
}
