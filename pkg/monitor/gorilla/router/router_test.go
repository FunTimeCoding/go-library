package router

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestRouter(t *testing.T) {
	assert.True(t, New() != nil)
}
