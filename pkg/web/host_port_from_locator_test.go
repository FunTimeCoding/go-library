package web

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestHostPortFromLocator(t *testing.T) {
	assert.String(
		t,
		"example.org:8080",
		HostPortFromLocator("https://example.org:8080/some/path"),
	)
}
