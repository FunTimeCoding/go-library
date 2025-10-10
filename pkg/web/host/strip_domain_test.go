package host

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestStripDomain(t *testing.T) {
	assert.String(t, "", StripDomain(""))
	assert.String(t, "org", StripDomain("org"))
	assert.String(t, "example.org", StripDomain("example.org"))
	assert.String(
		t,
		"test",
		StripDomain("test.example.org"),
	)
	assert.String(
		t,
		"test2",
		StripDomain("test2.test.example.org"),
	)
}
