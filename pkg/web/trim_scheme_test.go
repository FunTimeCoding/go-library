package web

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestTrimScheme(t *testing.T) {
	assert.String(t, "localhost", TrimScheme("https://localhost"))
	assert.String(t, "localhost", TrimScheme("http://localhost"))
}
