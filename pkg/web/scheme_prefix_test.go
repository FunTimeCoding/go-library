package web

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestSchemePrefix(t *testing.T) {
	assert.String(t, "https://", SchemePrefix(true))
	assert.String(t, "http://", SchemePrefix(false))
}
