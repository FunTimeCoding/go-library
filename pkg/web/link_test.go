package web

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestLink(t *testing.T) {
	assert.String(
		t,
		"https://localhost:9000",
		Link(Localhost, 9000, true),
	)
	assert.String(
		t,
		"https://localhost",
		Link(Localhost, 0, true),
	)
	assert.String(
		t,
		"http://localhost:9000",
		Link(Localhost, 9000, false),
	)
	assert.String(
		t,
		"http://localhost",
		Link(Localhost, 0, false),
	)
}
