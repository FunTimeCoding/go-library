package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestToIntegersStrict(t *testing.T) {
	assert.Any(t, []int{}, ToIntegersStrict([]string{}))
	assert.Any(t, []int{5}, ToIntegersStrict([]string{"5"}))
	assert.Any(t, []int{1, 2}, ToIntegersStrict([]string{"1", "2"}))
}
