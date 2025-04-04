package locator

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestRemoveDomain(t *testing.T) {
	assert.String(t, "", RemoveDomain(""))
	assert.String(t, "org", RemoveDomain("org"))
	assert.String(t, "example.org", RemoveDomain("example.org"))
	assert.String(
		t,
		"test",
		RemoveDomain("test.example.org"),
	)
	assert.String(
		t,
		"test2",
		RemoveDomain("test2.test.example.org"),
	)
}
