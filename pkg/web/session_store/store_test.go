package session_store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestStore(t *testing.T) {
	assert.NotNil(t, New())
}
