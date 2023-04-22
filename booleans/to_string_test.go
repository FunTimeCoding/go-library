package booleans

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestToString(t *testing.T) {
	assert.String(t, "1", ToString(true))
	assert.String(t, "0", ToString(false))
}
