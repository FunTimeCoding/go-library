package helper

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestHelper(t *testing.T) {
	assert.String(
		t,
		"http://localhost/example",
		ToWebLink("http://localhost/api/example"),
	)
}
