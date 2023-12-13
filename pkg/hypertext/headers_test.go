package hypertext

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestHeaders(t *testing.T) {
	assert.Any(
		t,
		[]string{
			"Example h1",
			"Example h2",
			"Example h3",
			"Example h4",
			"Example h5",
			"Example h6",
		},
		Headers(document(fixtureFile("test.html"))),
	)
}
