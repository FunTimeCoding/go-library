package repository

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestRepository(t *testing.T) {
	assert.True(t, New("", "") != nil)
}
