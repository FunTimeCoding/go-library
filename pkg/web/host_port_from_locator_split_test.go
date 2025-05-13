package web

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestHostPortFromLocatorSplit(t *testing.T) {
	host, port := HostPortFromLocatorSplit(
		"https://example.org:8080/some/path",
	)
	assert.String(t, "example.org", host)
	assert.Integer(t, 8080, port)
}
