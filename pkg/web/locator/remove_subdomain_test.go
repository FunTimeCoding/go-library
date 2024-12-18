package locator

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestRemoveSubdomain(t *testing.T) {
	assert.String(t, "", RemoveSubdomain(""))
	assert.String(t, "org", RemoveSubdomain("org"))
	assert.String(t, "example.org", RemoveSubdomain("example.org"))
	assert.String(
		t,
		"example.org",
		RemoveSubdomain("test.example.org"),
	)
	assert.String(
		t,
		"test.example.org",
		RemoveSubdomain("test2.test.example.org"),
	)
}
