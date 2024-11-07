package web

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestHostFromLocator(t *testing.T) {
	assert.String(
		t,
		"example.org",
		HostFromLocator("https://example.org:8080/some/path"),
	)
	assert.String(
		t,
		"example.org",
		HostFromLocator("https://example.org/some/path"),
	)
}
