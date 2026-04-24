package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestToIntegers64Strict(t *testing.T) {
	assert.Any(t, []int64{}, ToIntegers64Strict([]string{}))
	assert.Any(t, []int64{5}, ToIntegers64Strict([]string{"5"}))
	assert.Any(t, []int64{1, 2}, ToIntegers64Strict([]string{"1", "2"}))
}
