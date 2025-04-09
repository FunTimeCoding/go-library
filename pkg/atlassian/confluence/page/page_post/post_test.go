package page_post

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestPost(t *testing.T) {
	assert.True(
		t,
		New("", "", "", "") != nil,
	)
}
