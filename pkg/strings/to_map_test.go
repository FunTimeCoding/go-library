package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestToMap(t *testing.T) {
	assert.Any(
		t,
		map[string]string{"A": "a", "B": "b"},
		ToMap([]string{"A=a", "B=b"}, "="),
	)
}
