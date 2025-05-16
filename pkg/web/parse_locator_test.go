package web

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestParseLocator(t *testing.T) {
	actual := ParseLocator("https://example.org:8000")

	assert.String(t, "https://example.org:8000", actual.String())

	assert.String(t, "https", actual.Scheme)
	assert.String(t, "example.org:8000", actual.Host)
	assert.String(t, "8000", actual.Port())
	assert.String(t, "", actual.Path)
	assert.String(t, "", actual.RawQuery)
}
