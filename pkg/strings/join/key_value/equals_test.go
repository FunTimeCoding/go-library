package key_value

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestEquals(t *testing.T) {
	assert.String(t, "a=b", Equals("a", "b"))
}
