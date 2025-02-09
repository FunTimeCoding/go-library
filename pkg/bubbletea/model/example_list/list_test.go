package example_list

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestGroceryList(t *testing.T) {
	assert.True(t, New() != nil)
}
