package page_file

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestFile(t *testing.T) {
	assert.True(t, New("", "") != nil)
}
