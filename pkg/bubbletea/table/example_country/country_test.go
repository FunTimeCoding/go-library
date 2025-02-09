package example_country

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestCountry(t *testing.T) {
	assert.True(t, New() != nil)
}
