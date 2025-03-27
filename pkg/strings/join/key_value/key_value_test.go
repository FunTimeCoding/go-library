package key_value

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestEquals(t *testing.T) {
	assert.String(t, "a=b", Equals("a", "b"))
	assert.String(t, "a b", Space("a", "b"))
	assert.String(t, "ab", Empty("a", "b"))
	assert.String(t, "a:b", Colon("a", "b"))
}
