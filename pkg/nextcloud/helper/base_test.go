package helper

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestBase(t *testing.T) {
	assert.String(t, "https://example.org", Base("example.org"))
}
