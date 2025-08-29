package reader

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestReader(t *testing.T) {
	assert.NotNil(t, New())
}
