package maps

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestSortKeys(t *testing.T) {
	assert.Strings(
		t,
		[]string{"a", "b"},
		SortKeys(map[string]string{"b": "2", "a": "1"}),
	)
}
