package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestContainsLower(t *testing.T) {
	assert.True(t, ContainsLower("TestFixture", "fixture"))
	assert.False(t, ContainsLower("TestFixture", "other"))
}
