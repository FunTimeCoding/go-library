package strings

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestToIntegers(t *testing.T) {
	assert.Any(t, []int{}, ToIntegers([]string{}))
	assert.Any(t, []int{5}, ToIntegers([]string{"5"}))
	assert.Any(t, []int{1, 2}, ToIntegers([]string{"1", "2"}))
}
