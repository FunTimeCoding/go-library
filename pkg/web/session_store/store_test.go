package session_store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestStore(t *testing.T) {
	assert.True(t, New() != nil)
}
