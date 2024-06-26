package web

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestScheme(t *testing.T) {
	assert.String(t, "https", Scheme(true))
	assert.String(t, "http", Scheme(false))
}
