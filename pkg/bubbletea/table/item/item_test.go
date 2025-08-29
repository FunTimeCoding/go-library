package item

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestItem(t *testing.T) {
	assert.NotNil(t, New(true))
}
